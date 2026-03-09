package validator

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/pingcap/tidb/pkg/parser/ast"
	_ "github.com/pingcap/tidb/pkg/parser/test_driver"
	"gopkg.in/yaml.v3"
)

var targetStatements = []string{"select", "update", "delete"}

type sqlcConfig struct {
	SQL []struct {
		Queries string `yaml:"queries"`
	} `yaml:"sql"`
}

type Result struct {
	File      string
	Statement string
	Name      string
	Type      string
	Valid     bool
	Err       error
}

func (r Result) String() string {
	location := r.File
	if r.Name != "" {
		location = fmt.Sprintf("%s (%s)", r.File, r.Name)
	}
	if r.Err != nil {
		return fmt.Sprintf("ERROR %s: %v", location, r.Err)
	}
	if r.Valid {
		return fmt.Sprintf("OK %s [%s]", location, r.Type)
	}
	return fmt.Sprintf("NG %s [%s]: missing condition on schools.id", location, r.Type)
}

func ValidateSQLCProject(configPath string) ([]Result, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(configPath)
	var results []Result
	for _, block := range cfg.SQL {
		queryRoot := filepath.Join(baseDir, block.Queries)
		err := filepath.WalkDir(queryRoot, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if d.IsDir() || filepath.Ext(path) != ".sql" {
				return nil
			}
			fileResults, err := ValidateSQLFile(path)
			if err != nil {
				return err
			}
			results = append(results, fileResults...)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

func ValidateSQLFile(path string) ([]Result, error) {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	p := parser.New()
	stmtNodes, _, err := p.Parse(string(sqlBytes), "", "")
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	annotations := parseAnnotations(string(sqlBytes))
	results := make([]Result, 0, len(stmtNodes))
	for idx, stmt := range stmtNodes {
		stmtType, ok := statementType(stmt)
		if !ok {
			continue
		}

		name := ""
		if idx < len(annotations) {
			name = annotations[idx]
		}

		valid, err := validateStatement(stmt)
		results = append(results, Result{
			File:      path,
			Statement: restoreSQL(stmt),
			Name:      name,
			Type:      stmtType,
			Valid:     valid,
			Err:       err,
		})
	}

	return results, nil
}

func loadConfig(path string) (*sqlcConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var cfg sqlcConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return &cfg, nil
}

func parseAnnotations(sqlText string) []string {
	lines := strings.Split(sqlText, "\n")
	names := make([]string, 0)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "-- name:") {
			continue
		}
		fields := strings.Fields(trimmed)
		if len(fields) >= 3 {
			names = append(names, fields[2])
		}
	}
	return names
}

func statementType(stmt ast.StmtNode) (string, bool) {
	switch stmt.(type) {
	case *ast.SelectStmt:
		return "SELECT", true
	case *ast.UpdateStmt:
		return "UPDATE", true
	case *ast.DeleteStmt:
		return "DELETE", true
	default:
		return "", false
	}
}

func validateStatement(stmt ast.StmtNode) (bool, error) {
	tables := collectTables(stmt)
	if len(tables) == 0 {
		return false, fmt.Errorf("schools table not found in statement")
	}

	whereExpr := whereExpression(stmt)
	if whereExpr == nil {
		return false, nil
	}

	checker := schoolIDConditionChecker{
		targetAliases: tables,
	}
	whereExpr.Accept(&checker)
	return checker.found, nil
}

func whereExpression(stmt ast.StmtNode) ast.ExprNode {
	switch node := stmt.(type) {
	case *ast.SelectStmt:
		if node.Where == nil {
			return nil
		}
		return node.Where
	case *ast.UpdateStmt:
		return node.Where
	case *ast.DeleteStmt:
		return node.Where
	default:
		return nil
	}
}

func collectTables(stmt ast.StmtNode) []string {
	collector := tableCollector{
		aliases: make(map[string]struct{}),
	}
	stmt.Accept(&collector)

	aliases := make([]string, 0, len(collector.aliases))
	for alias := range collector.aliases {
		aliases = append(aliases, alias)
	}
	slices.Sort(aliases)
	return aliases
}

type tableCollector struct {
	aliases map[string]struct{}
}

func (c *tableCollector) Enter(in ast.Node) (ast.Node, bool) {
	source, ok := in.(*ast.TableSource)
	if !ok {
		return in, false
	}

	name, alias, matched := extractSchoolsAlias(source)
	if matched {
		c.aliases[name] = struct{}{}
		if alias != "" {
			c.aliases[alias] = struct{}{}
		}
	}

	return in, false
}

func (c *tableCollector) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extractSchoolsAlias(source *ast.TableSource) (string, string, bool) {
	tableName, ok := source.Source.(*ast.TableName)
	if !ok {
		return "", "", false
	}
	if !strings.EqualFold(tableName.Name.String(), "schools") {
		return "", "", false
	}

	alias := source.AsName.String()
	return tableName.Name.String(), alias, true
}

type schoolIDConditionChecker struct {
	targetAliases []string
	found         bool
}

func (c *schoolIDConditionChecker) Enter(in ast.Node) (ast.Node, bool) {
	if c.found {
		return in, true
	}

	expr, ok := in.(ast.ExprNode)
	if !ok {
		return in, false
	}
	if hasSchoolIDCondition(expr, c.targetAliases) {
		c.found = true
		return in, true
	}
	return in, false
}

func (c *schoolIDConditionChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func hasSchoolIDCondition(expr ast.ExprNode, targetAliases []string) bool {
	switch node := expr.(type) {
	case *ast.BinaryOperationExpr:
		if isSchoolIDColumn(node.L, targetAliases) || isSchoolIDColumn(node.R, targetAliases) {
			return true
		}
	case *ast.PatternInExpr:
		return isSchoolIDColumn(node.Expr, targetAliases)
	case *ast.BetweenExpr:
		return isSchoolIDColumn(node.Expr, targetAliases)
	case *ast.IsNullExpr:
		return isSchoolIDColumn(node.Expr, targetAliases)
	case *ast.PatternLikeOrIlikeExpr:
		return isSchoolIDColumn(node.Expr, targetAliases)
	case *ast.ParenthesesExpr:
		return hasSchoolIDCondition(node.Expr, targetAliases)
	}
	return false
}

func isSchoolIDColumn(expr ast.ExprNode, targetAliases []string) bool {
	columnExpr, ok := expr.(*ast.ColumnNameExpr)
	if !ok {
		return false
	}

	column := columnExpr.Name
	if !strings.EqualFold(column.Name.String(), "id") {
		return false
	}

	table := column.Table.String()
	if table == "" {
		return len(targetAliases) == 1
	}

	for _, alias := range targetAliases {
		if strings.EqualFold(alias, table) {
			return true
		}
	}
	return false
}

func restoreSQL(stmt ast.StmtNode) string {
	text := strings.TrimSpace(stmt.Text())
	return strings.ReplaceAll(text, "\n", " ")
}
