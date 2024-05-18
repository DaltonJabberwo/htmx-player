package main

type VideoData struct {
    Id int
    Framerate int
    Time int32
}

func newVideoData(id int, fps int, time int32) VideoData {
    return VideoData{
        Id: id,
        Framerate: fps,
        Time: time,
    }
}
