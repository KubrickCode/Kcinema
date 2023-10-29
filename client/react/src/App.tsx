import axios from "axios";
import { useState } from "react";

const App = () => {
  const [email, setEmail] = useState("");
  const [title, setTitle] = useState("");
  const [message, setMessage] = useState("");

  const getUsers = async () => {
    const response = await axios.get("http://localhost:3000/users");
    console.log(response);
  };

  const createUser = async () => {
    await axios.post("http://localhost:3000/user", { email });
  };

  const getBoards = async () => {
    const response = await axios.get("http://localhost:3000/boards");
    console.log(response);
  };

  const createBoard = async () => {
    await axios.post("http://localhost:3000/board", { title });
  };

  const getMessage = async () => {
    const response = await axios.get("http://localhost:3000/messages");
    console.log(response);
  };

  const createMessage = async () => {
    await axios.post("http://localhost:3000/message", { message });
  };

  const getGoUser = async () => {
    const response = await axios.get("http://localhost:5000/users");
    console.log(response);
  };

  const createGoUser = async () => {
    await axios.post("http://localhost:5000/user", { email });
  };

  return (
    <>
      <div>
        <button onClick={getUsers}>유저 불러오기</button>
        <input
          type="text"
          onChange={(e) => {
            setEmail(e.target.value);
          }}
          placeholder="이메일"
        ></input>
        <button onClick={createUser}>유저 저장</button>
      </div>
      <div>
        <button onClick={getBoards}>게시글 불러오기</button>
        <input
          type="text"
          onChange={(e) => {
            setTitle(e.target.value);
          }}
          placeholder="게시글 제목"
        ></input>
        <button onClick={createBoard}>게시글 저장</button>
      </div>
      <div>
        <button onClick={getMessage}>메시지 불러오기</button>
        <input
          type="text"
          onChange={(e) => {
            setMessage(e.target.value);
          }}
          placeholder="메시지 내용"
        ></input>
        <button onClick={createMessage}>메시지 저장</button>
      </div>
      <div>
        <button onClick={getGoUser}>go 유저 불러오기</button>
        <input
          type="text"
          onChange={(e) => {
            setEmail(e.target.value);
          }}
          placeholder="이메일"
        ></input>
        <button onClick={createGoUser}>go 유저 저장</button>
      </div>
    </>
  );
};

export default App;
