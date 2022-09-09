import React from 'react'
import CreatePost from './CreatePost'
import DisplayPost from './DisplayPost'

export default ()=>{
   // return (xml/hml);

    return(
        //<h1>Welcome to React</h1> //this is XML;
        // we cannot use same tag twice as it is not html; XML needs container like div blocks
       //<div> <h1>Welcome to React</h1>
       // <h1>Welcome to React</h1></div>
       <div>
           <div><CreatePost/></div>
           <div><DisplayPost/></div>
       </div>
    )
}