<!DOCTYPE html>

<html>
<head>
  <title>System Monitor</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.6/css/bootstrap.min.css" />
<link rel="stylesheet" href="/static/css/index.css" />
<body ng-app="system">
<div class="panel panel-info" ng-controller="sysinfoshow">
	<cpu id="cpu" date="date" data="data"></cpu>
</div>
</body>
<script src="//cdn.bootcss.com/jquery/2.2.4/jquery.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
<script src="//cdn.bootcss.com/angular.js/1.5.7/angular.min.js"></script>
<script src="/static/js/echarts.js"></script>
<script type="text/javascript">var pcid = {{.pcid}}</script>
<script src="/static/js/pcinfo.js"></script>
</html>