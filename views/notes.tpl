<head>
    <title>Notes</title>
</head>

<body>
    <header>
        <h1 class="logo">Welcome to Notes!</h1>
        <h2 class="description">
            This Notes is just for fun!
        </h2>
    </header>
    <div class="notes">
        
        <form action="http://127.0.0.1:8080/notes" method="POST">
        Title:<br>
        <input type="text" name="Title" value="Title">
        <br>Date:<br>
        <input type="text" name="Date" value="">
        <br>Data.Val:<br>
        <input type="text" name="Val" value="Content">
        <br>
        <input type="submit" value="Submit">
        </form> 



        <table border="1">
        <tr>
        <th>NoteId</th>
        <th>Title</th>
        <th>Date</th>
        <th>Data.Val</th>
        </tr>
        {{range $key ,$v:=.Notes}}
        <tr>
        <td>{{$key}}</td>
        <td>{{$v.Title}}</td>
        <td>{{$v.Date}}</td>
        <td>{{$v.Data.Val}}</td>
        </tr>
        {{end}}
        </table>
        
    </div>
<body>