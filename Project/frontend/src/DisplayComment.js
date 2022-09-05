import React,{useState,useEffect} from 'react'
import axios from 'axios'


export default ({postId})=>{

    const [comments,updateComments] = useState([])
  
     const loadComments = async () =>{
       const resp  = await axios.get(`http://localhost:4002/api/v1/blog/post/${postId}/comment`)
       updateComments(resp.data)
     }
    useEffect(()=>{
        loadComments();
    },[])

  //    console.log(comments)
   const liOfComments = comments?.map(c=>{
       return (
           <li key={c.commentId}>
              {c.message}
           </li>
       );
   })

   return(
    // liOfComments.length>0?
      <ol>
      {liOfComments}
      </ol>
      //   :
      // <ol>

      // </ol>
    );

}