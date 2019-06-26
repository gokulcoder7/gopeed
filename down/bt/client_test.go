package bt

import (
	"testing"
)

func TestClient_AddTorrent(t *testing.T) {
	client := NewClient()
	metaInfo, err := ParseFromFile("testdata/Game.of.Thrones.S08E05.720p.WEB.H264-MEMENTO.torrent")
	if err != nil {
		panic(err)
	}
	torrent := &Torrent{
		client:   client,
		MetaInfo: metaInfo,
	}
	torrent.MetaInfo.Announce = "udp://tracker.opentrackr.org:1337/announce"
	torrent.MetaInfo.AnnounceList = [][]string{}
	err = torrent.Tracker()
	if err != nil {
		panic(err)
		return
	}
}
