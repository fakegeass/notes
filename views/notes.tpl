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
        
        <form action='{{urlfor "NotesController.Post"}}' method="POST" enctype="application/x-www-form-urlencoded">
        <br>Title:<br>
        <input type="text" name="Title" placeholder="Title">
        <br>Date:<br>
        <input type="text" name="Date" placeholder="Date">
        <br>Val:<br>
        <input type="text" name="Date.Val" placeholder="content">
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