import React,{useState} from "react"; 
import axios from "axios"; //deals with post and get request to send data to a running service
export default () => {
    const labTitle="Title_Post"
    const [title,updateTitle] = useState("My First Post") 
   const onSubmitHandler = async (e) => {
      e.preventDefault();
      await axios.post("http://myproj.com/api/v1/blog/post",{
           title
       }).catch(e=>console.log(e.message))

updateTitle('')
   }
    return (
        <form onSubmit={onSubmitHandler}>
            <div className="form-group">
            <label>{labTitle}</label>
            <input type="text" className="form-control" value={title} onChange={(e) => updateTitle(e.target.value)}/> 
            </div>
            <button className="btn btn-primary">submit</button>
        <div>You entered is : {title}</div>
        </form>

    )
}