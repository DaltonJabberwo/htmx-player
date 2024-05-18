package main

type Player struct {
    VideoData VideoData
    BufferData BufferData
}

func newPlayer(id int, fps int, time int32) Player {
    return Player {
        VideoData: newVideoData(id, fps, time),
        BufferData: newBufferData(id, time),
    }
}

type BufferData struct {
    Id int
    EndBuffer int32
}

func newBufferData(id int, endBuffer int32) BufferData {
    return BufferData{
        Id: id,
        EndBuffer: endBuffer,
    }
}
