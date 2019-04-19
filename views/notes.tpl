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
        <br>Notes ID:<br>
        <input type="text" name="UUID" placeholder={{.Num}}>
        <br>Title:<br>
        <textarea rows="1" name="Title" cols="50"></textarea>
        <br>Val:<br>
        <textarea rows="8" name="Data.Val" cols="50"></textarea>
        <br>
        <input type="submit" value="Submit">
        
        
        <!--利用js中的encodeURIComponent进行转义处理-->
        <a title="请输入待删除的Notes ID"><input type="button" value="delete" onclick="javasrctpt:window.location.href='/notes/delete?UUID='+encodeURIComponent(this.form.UUID.value)"></a>
        
        </form>
        


        <h3>Notes:</h3>
        <table width="1500" border="1" style="table-layout: fixed; word-wrap:break-word;">  
        <tr>
        <th width="100">Num</th>
        <th width="180">Title</th>
        <th width="350">Date</th>
        <th>Content</th>
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