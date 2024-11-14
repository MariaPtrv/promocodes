import './App.css';
import { useState } from "react";
import Form from "./components/Form";
import Message from './components/Message';

function App() {
  const [message, setMessage] = useState('');

  function usePromocode(promocode) {
    if (promocode.length > 0) {
      setMessage("Промокод применяется ...");
      getServerSideProps(promocode);
    } else {
      setMessage("Промокод не может быть пустым 😯")
    }
  }

  async function getServerSideProps(promocode) {
    // const url = "http://localhost:3001/promocodes";
    const url = "http://localhost:8000/admin/promocodes/promocode/";

    try {
      const myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");
      const response = await fetch(url, {
        headers: myHeaders,
        method: "POST",
        body: JSON.stringify({     "promocode": "223232",
          "reward_id": 5,
          "max_uses": 4 }),
      });
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }

      const json = await response.json();
      console.log(json);
      
    } catch (error) {
      console.error(error.message);
    }


  }

  return (
    <div className="App">
      <div className="App-body">
        <p>
          Введите промокод
        </p>
        <Form usePromocode={usePromocode}></Form>
        <Message message={message}></Message>
      </div>

    </div>
  );
}

export default App;
