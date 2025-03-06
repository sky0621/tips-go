package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	components := createComponents()

	for _, cmp := range components {
		wg.Add(1)
		go cmp.exec(wg)
	}

	wg.Wait()
}

func createComponents() []component {
	// けっこうエグい構造を作っておく
	return []component{
		newComposite("A", []component{
			newComposite("A-a", []component{
				newLeaf("A-a-1"),
			}),
			newComposite("A-b", []component{
				newLeaf("A-b-1"),
				newLeaf("A-b-2"),
			}),
		}),
		newComposite("B", []component{
			newLeaf("B-a"),
			newComposite("B-b", []component{
				newComposite("B-b-1", []component{
					newLeaf("B-b-1-1"),
					newLeaf("B-b-1-2"),
					newComposite("B-b-1-3", []component{
						newLeaf("B-b-1-3-1"),
					}),
				}),
			}),
			newLeaf("B-c"),
		}),
		newLeaf("C"),
		newComposite("D", []component{
			newComposite("D-d", []component{
				newLeaf("D-d-1"),
				newLeaf("D-d-2"),
				newLeaf("D-d-3"),
				newLeaf("D-d-4"),
			}),
		}),
	}
}

// ----------------------------------------------
// コンポジット構造
// ----------------------------------------------

type component interface {
	exec(wg *sync.WaitGroup)
}

func newComposite(name string, components []component) component {
	c := &composite{name: name}
	for _, cmp := range components {
		c.add(cmp)
	}
	return c
}

type composite struct {
	name     string
	children []component
}

func (c *composite) exec(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("[COMPOSITE]%s\n", c.name)

	wg2 := &sync.WaitGroup{}
	for _, child := range c.children {
		wg2.Add(1)
		go child.exec(wg2)
	}
	wg2.Wait()
}

func (c *composite) add(cmp component) {
	c.children = append(c.children, cmp)
}

func newLeaf(name string) component {
	return &leaf{name: name}
}

type leaf struct {
	name string
}

func (l *leaf) exec(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("[LEAF]%s\n", l.name)
}
