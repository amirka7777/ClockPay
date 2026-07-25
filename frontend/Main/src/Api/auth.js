const API_URL = 'http://localhost:8080/api/auth/register'

async function register(username, password, confirm_password) {
  const response = await fetch(API_URL, {
    method: 'POST',

    headers: {
      'content-type': 'application/json'
    },
    body: JSON.stringify ({
      username: username,
      password: password,
      confirm_password: confirm_password,
    })
  });

  if (!response.ok) {
    const errorData = await response.json()
    alert(JSON.stringify(errorData)); 
    alert(errorData.message || '🔴Ошибка регистрации🔴')
  } else {
    const sucsessData = await response.json()
    return sucsessData;
  }
}

export default register;