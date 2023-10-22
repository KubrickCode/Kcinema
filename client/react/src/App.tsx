import axios from "axios";
import { useState } from "react";

const App = () => {
  const [email, setEmail] = useState("")

  const getUser = async() => {
    const response = await axios.get("http://localhost:3000/users");
    console.log(response)
  }

  const handleSubmit = async() => {
    await axios.post("http://localhost:3000/user",{email});
  }
  
  return(
  <>
  <button onClick={getUser}>유저 불러오기</button>
  <input type="text" onChange={(e)=>{setEmail(e.target.value)}}></input>
  <button onClick={handleSubmit}>유저 저장</button>
  </>
  )
}

export default App;
