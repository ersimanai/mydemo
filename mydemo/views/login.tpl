<html>
<head>
	<title>登录</title>
	<script type="text/javascript" src = "/static/js/test.js"></script>
</head>
<body>
	<form action="/login/login"  onsubmit = "checkpost()" method="post" >
    	用户名：<br>
    	<input name="username" type="text"><br/>
    	密码：<br>
    	<input name="password" type="text"><br/>
    
    	<input type="submit"  value="登录" ><br/>
	</form>
</body>
</html>