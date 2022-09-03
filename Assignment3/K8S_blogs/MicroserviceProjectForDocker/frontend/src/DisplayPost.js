import React,{useState,useEffect} from "react"; //useeffect for calling method at time of load and only once
//usestate for hooks on everyupdate
import axios from "axios";
import CreateComment from "./CreateComment";
import DisplayComment from "./DisplayComment";
export default () => {
 const [posts,updatePosts] = useState({}) // creating a hook for updates
const loadPost = async ()=>{
   const resp = await axios.get("http://myproj.com/api/v1/blog/query/post").catch(e=>console.log(e.message))
updatePosts(resp.data)
}

useEffect(() => {
    loadPost();
},[])//useeffect for call method at time of load and only once

//creating dynmic number of cards
const cardofpost = Object.values(posts).map(p=>{
    return (
        <div className="card" style={{ width:"30%", marginBottom:"20px"}}>

<div className="card-body" key={p.id}>
    {p.title}

</div>
<div>
    <DisplayComment comments={p.comments}/>
    <CreateComment postId={p.id}/>
</div>
        </div>
    )
})

    return(
   // <div><h1>Display</h1></div> //do this initially
    <div className="d-flex felx-row flex-wrap justify-content-between">
{cardofpost}
    </div>


    )//creating dynmic number of cards
}