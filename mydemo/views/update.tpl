<html>
<head>
	<title>修改</title>
</head>
<body>
	<form action="/update/update"  onsubmit = "checkpost()" method="post">
    	用户名：<br>
    	<input name="username" type="text"><br/>
    	旧密码：<br>
    	<input name="oldpassword" type="text"><br/>
    	新密码：<br>
    	<input name="newpassword" type="text"><br/>
    	<input type="submit"  value="修改" ><br/>
	</form>
</body>
</html>