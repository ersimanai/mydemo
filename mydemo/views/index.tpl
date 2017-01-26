<!DOCTYPE html>

<html>
<head>
  <title>demo</title>  
</head>

<body>
  <header>
    <h1 class="logo">Welcome to mydemo</h1>
  </header>
  <form action="/{{.register}}"  onsubmit = "checkpost()" method="post" {{.register}}>
    用户名：<br>
    <input name="username" type="text"><br/>
    密码：<br>
    <input name="password" type="text"><br/>
    
    <input type="submit"  value="注册" ><br/>

  </form>
<script>
function checkpost(){
    alert("lalalala")
  }
</script>
 </body>
</html>
