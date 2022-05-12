import React,{ Component } from 'react'
import VideoPlayer from './Player'


class VideoPage extends Component {
      
     
    render() {
        let videoJsOptions = {
            autoplay: true,
            controls: true,
            sources: [{
              src: 'rtmp://192.168.208.1:9935/video/tt2.mp4'
            }]
          }
        return(
          <div >
              <VideoPlayer { ...videoJsOptions } />
              <div>测试测试</div>
          </div>
        )
    }
}
  
export default VideoPage

