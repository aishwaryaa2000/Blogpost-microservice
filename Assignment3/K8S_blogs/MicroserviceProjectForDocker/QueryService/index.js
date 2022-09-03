const express = require('express')
const axios = require('axios')
const bodyParser = require('body-parser')
const cors = require('cors')

const app=express();
app.use(bodyParser.json());
app.use(cors())

const posts={} // this wil contain all posts; posts will get populated based on eventbus and event fires type.
// app.get('/api/v1/blog/post',(req,resp)=>{
app.get('/api/v1/blog/query/post',(req,resp)=>{

    resp.send(posts);
})

const handleEvent = (type,data) => {
    //for posts
    if (type == "Post Created"){
        const {id,title} = data;
        posts[id] = {id,title,comments:[]}
        return;
    }
    //for comments
    if (type == "Comment Created"){
        const {postId,commentId,message} = data;
        const post = posts[postId]
        post.comments.push({commentId,message})
        return;
    }
}
//this is imp for queryservice as on this ; post will be populated.
app.post("/eventbus/event/listener",(req,resp)=>{
    const {type,data} = req.body
    console.log("Recived Event",type)
    handleEvent(type,data);
 
    resp.send({});
 });

 app.listen(4003,async ()=>{
     //get all events that have occured
     const resp = await axios.get('http://eventbus-serv:4005/eventbus/event').catch(e=>console.log(e.message));
     if(resp.data!=null){
        const events = resp.data;
        for(let e of events){
         handleEvent(e.type,e.data);
        }
        console.log("4003  okk ttttt testing")
     }
     console.log("4003 running okk testing")

 })