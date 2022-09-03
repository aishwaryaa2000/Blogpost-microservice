const express = require('express')
const cors  = require('cors')
const  bodyParser = require('body-parser')
const uuid = require('uuid')
const axios = require('axios')

const app = express();
app.use(cors())
app.use(bodyParser.json())

const postsWithComments={}
app.get('/api/v1/blog/post/:postId/comment',(req,resp)=>{
  const postId = req.params.postId
  const comments = postsWithComments[postId]||[]
  resp.send(comments)
});

app.post('/api/v1/blog/post/:postId/comment',async (req,resp)=>{
    const commentId = uuid.v4();    
    const {message}   = req.body; // {message:'',likes:10,foo:'bar'}
    const postId = req.params.postId;
    const comments = postsWithComments[postId] || []
    comments.push({commentId,message})
    postsWithComments[postId] = comments
         //event after successfull addition of Blogpost, we need to send an event to eventbus.
         await axios.post("http://eventbus-serv:4005/eventbus/event",{
          type:"Comment Created",
          data:{postId,commentId,message}
       }).catch(e=>console.log(e.message))

    resp.status(201).send({commentId,message})
});
app.post("/eventbus/event/listener",(req,resp)=>{
  const {type} = req.body
  console.log("Recived Event",type)
  resp.send({});
});

app.listen(4002,()=>{
   console.log("comments service listening at 4002")

})
