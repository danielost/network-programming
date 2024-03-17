const apiUrl = 'http://localhost:8000';

function sendData() {
    let message = document.getElementById("message").value
    let result = document.getElementById("result")

    const msg = {
        message: message,
    };

    fetch(`${apiUrl}/sort_words`, {
        headers: {
            "Content-Type": "application/json"
        },
        method: "POST",
        body: JSON.stringify(msg),
    })
        .then(response => {
            if (!response.ok) {
                alert('Помилка! Статус ' + response.statusText)
                return
            }

            response.json().then(data => {
                result.innerHTML = 'Результат: ' + data.message
            })
        })
}
