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
    const url = "http://localhost:8001/promocodes/promocode/use";

    try {
      const myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");
      const response = await fetch(url, {
        headers: myHeaders,
        method: "POST",
        body: JSON.stringify({
          "promocode": promocode,
          "user_id": 123,
        }),
      });
      if (!response.ok) {
        setMessage("Промокод не применен 😒 Проверьте и попробуйте еще раз.")
        throw new Error(`Response status: ${response.status}`);
      }

      const json = await response.json();
      console.log(json);

     if (json.id) {
          let msg = "Промокод применен 😊"
          msg += json.description.length > 0 ? " Промокод " + json.description.toLowerCase() : "."
        setMessage(msg)
      } else {
        switch (json.status) {
          case 0:
            setMessage("Промокод уже был применен ранеe 😊")
            break;
          case 1:
            setMessage("Сроки применения промокода истекли ⌛️")
            break;
          case 2:
            setMessage("Промокод больше не применяется.")
            break;
          default:
            setMessage("Промокод не применен 😒 Проверьте и попробуйте еще раз.")
        }
      }

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
