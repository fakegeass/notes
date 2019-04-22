<head><title>file</title></head>

<body>
    <h1>File Share:</h1>
    <div><form id="upload-form" action='/upload' method="post" enctype="multipart/form-data" >
    <input type="file" id="upload" name="upload" /> 
    <input type="submit" value="Upload" />
    </form></div>
    
    <table>
    <tr>
    <th align="left" width="100">File Name</th>
    <th align="left" width="180">Size</th>
    <th align="left" width="180">Date</th>
    </tr>
    {{range $k,$v :=.file}}
    <tr>
    <td><a href="/file/{{$v.Name}}">{{$v.Name}}</td>
    <td>{{$v|isFile}}</td>
    <td>{{$v.ModTime|format}}</td>
    </tr>
    {{end}}
    </table>
</body>