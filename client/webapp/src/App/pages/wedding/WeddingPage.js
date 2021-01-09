import React,{ Component } from 'react'
import {Image,message, Result} from "antd"
import {get} from '../../utils/RequestREST'

class WeddingPage extends Component {
    state={images:[]}
    render() {
        return(
          <div >
            <Image.PreviewGroup className="imageGroup">
              {this.state.images?this.state.images.map((item,index)=>{
                  return <Image width={600}  src={"http://127.0.0.1/"+item.uri}/>
                }):null}
            </Image.PreviewGroup>
          </div>
        )
    }
    componentDidMount(){
        //get("/rest/assets/images").then(res=>res.json()).then(res=>{
        //    this.setState({images:res})
        //})
    }
}
  
export default WeddingPage