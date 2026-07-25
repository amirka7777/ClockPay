import React from 'react';
import { useState } from 'react';
import register from '../Api/auth'

const Register = () => {
    const [UserName, setUserName] = useState('')
    const [Password, setPassword] = useState('')
    const [ConfirmPassword, setConfirmPassword] = useState('')

    async function handleSubmit(e) {
        e.preventDefault()

        if (Password === ConfirmPassword) {
            try {
                const response = await register(UserName, Password, ConfirmPassword);
                console.log('Успешно:', response);
                window.location.href = '/'; 
            } catch (error) {
                console.error('Ошибка регистрации:', error);
                alert(error.message || 'Ошибка регистрации');
            };
        };

    };

    function togglePasswordVisibility(inputId, checkboxId) {
        const input = document.getElementById(inputId);
        const checkbox = document.getElementById(checkboxId);

        if (input && checkbox) {
        if (checkbox.checked) {
            input.type = 'text';   
            } else {
            input.type = 'password'; 
            }
        }
    }

    return (
        <div className="Registration">

            <div className='backdrop'>

                <form onSubmit={handleSubmit}>

                    <h1>Регистрация</h1>

                        <input type='text'
                               placeholder='Введите ник'
                               value={UserName}
                               onChange={(e) => setUserName(e.target.value)}/>

                    <div className='check-password'>
                        <input type='password'
                               placeholder='Введите пароль'
                               value={Password}
                               onChange={(e) => setPassword(e.target.value)}
                               id="password-field"/>

                        <input id="password-checkbox" 
                               type="checkbox" 
                               onChange={() => togglePasswordVisibility('password-field', 'password-checkbox')}/>
                    </div>

                    <div className='check-password'>
                        <input type='password'
                               placeholder='Подтвердите пароль'
                               value={ConfirmPassword}
                               onChange={(e) => setConfirmPassword(e.target.value)}
                               id="confirm-password-field"/>

                        <input id="confirm-password-checkbox" 
                               type="checkbox" 
                               onChange={() => togglePasswordVisibility('confirm-password-field', 'confirm-password-checkbox')}/>
                    </div>

                    <button type='submit'>Зарегистрироваться</button>

                    <p>Уже есть аккаунт? 
                        <a href='/Login'>войти</a>
                    </p>
                </form>
            </div>
        </div>
    );
}

export default Register;