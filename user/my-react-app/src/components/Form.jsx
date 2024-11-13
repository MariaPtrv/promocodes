import { useState } from "react";

function Form(props) {
    const [promocode, setPromocode] = useState('');


    function handleSubmit(event) {
        event.preventDefault();
        props.usePromocode(promocode);
    }

    function handleChange(event) {
        setPromocode(event.target.value);
    }

    return (
        <form onSubmit={handleSubmit}>
            <div className="promocode-input">
                <input type="text" id="promocode" autoComplete="off"
                    value={promocode}
                    onChange={handleChange}
                ></input>
                <button type="submit">Применить</button>
            </div>
        </form>
    );
}

export default Form;