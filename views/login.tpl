<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
</head>
<body>
    <form action='/login' method="post" enctype="multipart/form-data">
        <div class="field-content">
            User Name：<input name="user" type="text" />
        </div>
        <div class="field-content">
            Password：<input name="pass" type="password" />
        </div>
        <div class="field-content">
            <input type="submit" value="登陆" />
        </div>
    </form>
</body>