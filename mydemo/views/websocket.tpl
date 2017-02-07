<html>
    <head>
        <meta charset="utf-8" />
        <title>聊天室</title>
       
         </head>
    <body >
     
        <div class="container">
        <h3>用户名: <span id="uname" > {{.UserName}}</span></h3>
            <h4>发送框:</h4>
            <form class="form-inline">
                <div class="col-md-6 form-group">
                    <input id="sendbox" type="text" class="form-control" onkeydown="if(event.keyCode==13)return false;" required>
                </div>
                <button id="sendbtn" type="button" class="btn btn-default">{{.btn}}</button>
            </form>
        </div>

        <div class="container">
            <h3>聊天记录:</h3>
            <ul id="chatbox">
                 <li> 欢迎加入聊天室 </li>
            </ul>
        </div>

         <script type="text/javascript" src = "http://cdn.static.runoob.com/libs/jquery/1.10.2/jquery.min.js">
         </script>
         <script type="text/javascript" src="/static/js/websocket.js"></script>
    </body>
</html>