package demo

import (
	"embed"
	"errors"
	"math/rand"
	"path/filepath"
	"strings"
)

const dirName = "assets"

//go:embed assets
var assets embed.FS

type Resource struct {
	Name  string
	Text  string
	Audio []byte
}

type ResourcePool struct {
	m    map[string]Resource
	list []Resource
}

func newResourcePool(list []Resource) *ResourcePool {
	m := make(map[string]Resource, len(list))
	for _, resource := range list {
		m[resource.Text] = resource
	}
	return &ResourcePool{
		m:    m,
		list: list,
	}
}

func NewResourcePool() (*ResourcePool, error) {
	entries, err := assets.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	var res []Resource
	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dirName, entry.Name())
			subEntries, err := assets.ReadDir(subDir)
			if err != nil {
				return nil, err
			}
			r := Resource{
				Name: entry.Name(),
			}
			for _, subEntry := range subEntries {
				fPath := filepath.Join(subDir, subEntry.Name())
				if subEntry.Name() == "text.txt" {
					text, err := assets.ReadFile(fPath)
					if err != nil {
						return nil, err
					}
					r.Text = string(text)
				} else if strings.HasPrefix(subEntry.Name(), "audio") {
					audio, err := assets.ReadFile(fPath)
					if err != nil {
						return nil, err
					}
					r.Audio = audio
				} else {
					return nil, errors.New("unknown resource: " + subEntry.Name())
				}
			}
			res = append(res, r)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("no resource was found")
	}
	return newResourcePool(res), nil
}

func (p ResourcePool) RandomResource() *Resource {
	return &p.list[rand.Intn(len(p.list))]
}

func (p ResourcePool) FindAudioByTextOrRandom(text string) *Resource {
	r, ok := p.m[text]
	if ok {
		return &r
	}
	return p.RandomResource()
}

func (p ResourcePool) List() []Resource {
	return p.list[:]
}

func (p ResourcePool) Len() int {
	return len(p.list)
}
