import React,{useState} from "react";
import axios from "axios";
export default ({postId})=>{
   const [message,updateMessage]= useState("Message for comment 1")
   const handleMySubmit= async(e) =>{
       e.preventDefault();
        await axios.post(`http://localhost:4002/api/v1/blog/post/${postId}/comment`,{message}).catch(e=>console.log(e.message))
        updateMessage('')
   }

   return(
        <form onSubmit={handleMySubmit}>
            <div>
            <label className="form-group">Comment here</label>
            <input type="text" value={message} onChange={(e)=>updateMessage(e.target.value)} className="form-control"/>
            </div> {postId}<br/>
            <button className="btn btn-primary">submit</button>
           
        </form>
    )
}