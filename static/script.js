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
                // トークンをローカルストレージに保存
                localStorage.setItem('auth_token', data.token);

                // トークンをヘッダーにセットしてmain.htmlを要求
                fetch('/main', {
                    method: 'GET',
                    headers: {
                        'Authorization': `Bearer ${data.token}` // トークンをヘッダーに追加
                    }
                })
                    .then(response => response.text())  // レスポンスをHTMLとして受け取る
                    .then(html => {
                        // サーバーから返されたHTMLをbodyに挿入
                        document.body.innerHTML = html;
                    })
                    .catch(error => {
                        console.error('Error fetching main.html:', error);
                    });
            } else {
                alert('ログインに失敗しました');
            }
        });
});

