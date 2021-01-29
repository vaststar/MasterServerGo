import React,{ Component } from 'react'
import {Image,} from "antd"
import {get} from '../../utils/RequestREST'

class WeddingPage extends Component {
    state={images:[]}
    render() {
        return(
          <div >
            <Image.PreviewGroup className="imageGroup">
              {this.state.images?this.state.images.map((item,index)=>{
                  return <Image width={500}  src={item.uri}/>
                }):null}
            </Image.PreviewGroup>
          </div>
        )
    }
    componentDidMount(){
        get("/rest/assets/images").then(res=>res.json()).then(res=>{
            this.setState({images:res.data})
        })
    }
}
  
export default WeddingPage