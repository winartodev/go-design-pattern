package musicplaylist

import "fmt"

type IPlaylist interface {
	Clone() IPlaylist
}

type RegularPlaylist struct {
	Name  string
	Track []string
}

func (r *RegularPlaylist) Clone() IPlaylist {
	copiedTracks := make([]string, len(r.Track))
	for i := 0; i < len(r.Track); i++ {
		copiedTracks = append(copiedTracks, fmt.Sprintf("%s_clone", r.Track[i]))
	}

	return &RegularPlaylist{
		Name:  fmt.Sprintf("%s_clone", r.Name),
		Track: copiedTracks,
	}
}

type SmartPlaylist struct {
	Name  string
	Genre string
}

func (s *SmartPlaylist) Clone() IPlaylist {
	return &SmartPlaylist{
		Name:  fmt.Sprintf("%s_clone", s.Name),
		Genre: fmt.Sprintf("%s_clone", s.Genre),
	}
}

type PlaylistRegistry struct {
	prototypes map[string]IPlaylist
}

func NewPlaylistRegistry() *PlaylistRegistry {
	return &PlaylistRegistry{
		prototypes: make(map[string]IPlaylist),
	}
}

func (r *PlaylistRegistry) RegisterPrototype(name string, playlist IPlaylist) {
	r.prototypes[name] = playlist
}

func (r *PlaylistRegistry) UnregisterPrototype(name string) {
	delete(r.prototypes, name)
}

func (r *PlaylistRegistry) Clone(name string) IPlaylist {
	prototype, ok := r.prototypes[name]
	if ok {
		return prototype.Clone()
	}
	return nil
}
