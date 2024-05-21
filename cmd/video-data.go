package main

type VideoData struct {
    Id int
    Time int32
}

func newVideoData(id int, time int32) VideoData {
    return VideoData{
        Id: id,
        Time: time,
    }
}
