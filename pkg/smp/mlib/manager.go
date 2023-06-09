package library

import (
	"errors"
	"fmt"
)

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Write() {
	wirtejson(m.musics)
}

func (m *MusicManager) Read() {
	ele := readjson()
	for index, value := range ele {
		m.musics = append(m.musics, value)
		fmt.Println(index)
	}
	//
}

func (m *MusicManager) Add(music *MusicEntry) {

	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	// Remove the found item from the slice.
	if index < len(m.musics)-1 { // Element between first and last
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { // empty it.
		m.musics = make([]MusicEntry, 0)
	} else { // The last element
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}
