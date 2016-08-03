<!DOCTYPE html>

<html>
<head>
  <title>Spider Monitor System</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.6/css/bootstrap.min.css" />
<link rel="stylesheet" href="/static/css/index.css" />
<body ng-app="monitor">
<div class="panel panel-info" ng-controller="computers">
	<div class="panel-heading">
	  	<div class="panel-title">
				<a href="#">
				Spider Monitor System
				</a>
			<div style="float:right;">
			<smss></smss>
			</div>
		</div>
	</div>
	<div class="panel-body">
		<div ng-show="showDetails">
			<setting></setting>
		</div>
		<computer ng-repeat="computer in computers"></computer>
	</div>
</div>
</body>
<script src="//cdn.bootcss.com/jquery/2.2.4/jquery.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
<script src="//cdn.bootcss.com/angular.js/1.5.7/angular.min.js"></script>
<script type="text/javascript">var WebSocketIP = {{.ip}}</script>
<script src="/static/js/index.js"></script>
</html>
