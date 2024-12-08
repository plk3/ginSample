document.getElementById('register-form').addEventListener('submit', function (event) {
    event.preventDefault();

    const username = document.getElementById('register-username').value;
    const password = document.getElementById('register-password').value;

    fetch('/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    })
        .then(response => response.json())
        .then(data => {
            if (data.message) {
                alert('登録成功!');
            } else {
                alert('登録に失敗しました');
            }
        });
});

document.getElementById('login-form').addEventListener('submit', function (event) {
    event.preventDefault();

    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;

    fetch('/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                alert('ログイン成功!');
            } else {
                alert('ログインに失敗しました');
            }
        });
});
