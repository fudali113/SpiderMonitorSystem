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
	<div class="panel-heading">
	  	<div class="panel-title">
				<a href="#">
				System Monitor
				</a>
			<div style="float:right;">
				<div style="float:right;text-align:center;background:#53FF53" ng-style="SORBackground" class="ellipse" ng-click="switchStopRun()">
					<font color="white">--</font>
					<font ng-bind="stopOrRun"></font>
					<font color="white">--</font>
				</div>
			</div>
		</div>
	</div>
	<div class="panel-body">
		<cpu id="cpu" date="date" cpudata="cpudata" memdata="memdata"></cpu>
    <div class="panel panel-info" style="height:auto;">
    	<div class="panel-body">
        <font size="5">process info</font><font color="white">--------</font><font size="5">total:</font><font size="5" ng-bind="procs.length"></font>
        <div style="float:right;text-align:center;background:white;" class="ellipse">
          <div class="form-group">
            <div class="col-sm-12">
              <select ng-model="procs_order" class="form-control">
                <option value="pid">pid</option>
                <option value="memper">mem</option>
                <option value="name">name</option>
                <option value="io">io</option>
                <option value="threads">threds</option>
              </select>
            </div>
          </div>
        </div>
        <div style="float:right;text-align:center;background:white;" class="ellipse">
          <font size="5">order by:</font>
        </div>
        <div style="float:right;text-align:center;background:white;" class="ellipse">
          <div class="form-group">
            <div class="col-sm-12">
              <input ng-model="procs_query" class="form-control" id="lastname" placeholder="input filter charset">
            </div>
          </div>
        </div>
        <div style="float:right;text-align:center;background:white;" class="ellipse">
          <font size="5">search:</font>
        </div>
    		<table class="table table-bordered table-hover " style="boder:1px">
           <thead>
              <tr>
                 <th>pid</th>
                 <th>name</th>
                 <th>threads</th>
                 <th>mem percent</th>
                 <th>io</th>
              </tr>
           </thead>
           <tbody>
             <tr ng-repeat="proc in procs | filter:procs_query | orderBy:procs_order" ng-click="showModal(proc)">
              <td ng-bind="proc.pid"></td>
              <td ng-bind="proc.name"></td>
              <td ng-bind="proc.threads"></td>
              <td ng-bind="proc.memper"></td>
              <td ng-bind="proc.io"></td>
             </tr>
           </tbody>
    		</table>
      </div>
    </div>
	</div>
  <!-- 模态框（Modal） -->
  <div class="modal fade" id="myModal" tabindex="-1" role="dialog"
     aria-labelledby="myModalLabel" aria-hidden="true">
     <div class="modal-dialog">
        <div class="modal-content">
           <div class="modal-header">
              <button type="button" class="close"
                 data-dismiss="modal" aria-hidden="true">
                    &times;
              </button>
              <h4 class="modal-title" id="myModalLabel">
                 <font size="5">process detail info</font>
              </h4>
           </div>
           <div class="modal-body">
              <p ng-bind="modalContent.pid"></p>
           </div>
        </div><!-- /.modal-content -->
  </div><!-- /.modal -->
</div>
</body>
<script src="//cdn.bootcss.com/jquery/2.2.4/jquery.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
<script src="//cdn.bootcss.com/angular.js/1.5.7/angular.min.js"></script>
<script src="/static/js/echarts.js"></script>
<script type="text/javascript">var pcid = {{.pcid}}</script>
<script src="/static/js/pcinfo.js"></script>
</html>
