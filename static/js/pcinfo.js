var system = angular.module('system', []);

system.directive('cpumemper', function() {
    return {
		scope:{
			id:"@",
			date:"=",
			cpudata:"=",
			memdata:"="
			},
        restrict: 'E',
        template: '<div style="height:300px;width:600px"></div>',
        replace: true,
		link: function($scope, element, attrs) {
			var getOption = function(){
				return {
				    title: {
				        text: 'cpu,mem占用百分比'
				    },
				    tooltip: {
				        trigger: 'axis'
				    },
				    legend: {
				        data:['cpu','mem']
				    },
				    grid: {
				        left: '3%',
				        right: '4%',
				        bottom: '3%',
				        containLabel: true
				    },
				    toolbox: {
				        feature: {
				            saveAsImage: {}
				        }
				    },
				    xAxis: {
				        type: 'category',
				        boundaryGap: false,
				        data: $scope.date
				    },
				    yAxis: {
				        type: 'value'
				    },
				    series: [
				        {
				            name:'cpu',
				            type:'line',
				            stack: '利用率',
				            data:$scope.cpudata
				        },
				        {
				            name:'mem',
				            type:'line',
				            stack: '占用率',
				            data:$scope.memdata
				        }
				    ]
				};
			}

			var dom = echarts.init(document.getElementById($scope.id));
			dom.setOption(getOption())
			$scope.$on( 'sysinfo.update.link', function( event ) {
				dom.setOption(getOption())
		    });
		}
    };
});

system.directive('cpu', function() {
	return {
		restrict: 'E',
		template: '<div><font size="5">User:{{cpu.user}}</font><br/>'+
		'<font size="5">System:{{cpu.system}}</font><br/>'+
		'<font size="5">Idle:{{cpu.idle}}</font><br/>'+
		'</div>',
		replace: true,
		link:function () {

		}
	}
});

system.directive('mem', function() {
	return {
		scope:{
			mem:"="
		},
		restrict: 'E',
		template: '<div><font size="5">Total:{{detailInfo.mem.total}}</font><br/>'+
		'<font size="5">Available:{{detailInfo.mem.available}}</font><br/>'+
		'<font size="5">Used:{{detailInfo.mem.used}}</font><br/>'+
		'<font size="5">UsedPercent:{{detailInfo.mem.usedPercent}}</font><br/>'+
		'</div>',
		replace: true,
		link:function () {

		}
	}
});

system.service( 'sysinfo', [ '$rootScope','$http', function( $rootScope,$http ) {
	$rootScope.updateTime = 5000
	$rootScope.stopOrRun = true

	var getNowTimeStr = function(){
		var d = new Date()
		var getDouble = function(i){
			return i < 10 ? '0'+i : i
		}
		return getDouble(d.getHours())+':'+getDouble(d.getMinutes())+':'+getDouble(d.getSeconds())
	}

	var service = {
		data:{
			date:[getNowTimeStr()],
			cpudata:[0],
			memdata:[50],
			detailInfo:{}
		},
		addData:function(data){
			this.data.date.push(getNowTimeStr())
			this.data.cpudata.push(data.cpu[0])
			this.data.memdata.push(data.memInfo.usedPercent)
      		this.data.SDI = data
			this.data.detailInfo.cpus = data.cpuInfo
			this.data.detailInfo.mem = jsonStringify(data.memInfo,null,2)
			this.data.detailInfo.io = jsonStringify(data.ioInfo," ")
			this.data.detailInfo.net = jsonStringify(data.netInfo,"")
			$rootScope.$broadcast('sysinfo.update');
		},
		addCpuMemData:function (data) {
			this.data.date.push(getNowTimeStr())
			this.data.cpudata.push(data.cpu)
			this.data.memdata.push(data.mem)
			$rootScope.$broadcast('sysinfo.update');
		}
	}

	function jsonStringify(data,space){
	    var seen=[];
	    return JSON.stringify(data,function(key,val){
	        if(!val||typeof val !=='object'){
	            return val;
	        }
	        if(seen.indexOf(val)!==-1){
	            return '[Circular]';
	        }
	        seen.push(val);
	        return val;
	    },space);
	}

	var getSysinfo = function(){
		if (!$rootScope.stopOrRun) return
		$http({
			url:'/'+pcid+'/info/all',
			method:'get',
		}).success(function(data){
			service.addData(data)
		})
	}
	var getSimpleInfo = function(){
		if (!$rootScope.stopOrRun) return
		$http({
			url:'/'+pcid+'/info/simple',
			method:'get',
		}).success(function(data){
			service.addCpuMemData(data)
		})
	}
	getSysinfo()
	getSimpleInfo()
	setInterval(getSysinfo,60000)
	setInterval(getSimpleInfo,5000)
	return service
}]);

system.controller('sysinfoshow',['$rootScope','$scope','$http','sysinfo',function($rootScope,$scope,$http,sysinfo){
	$scope.$on( 'sysinfo.update', function( event ) {
    $scope.date = sysinfo.data.date
    $scope.SDI = sysinfo.data.SDI
	$scope.cpudata = sysinfo.data.cpudata
	$scope.memdata = sysinfo.data.memdata
	$scope.detailInfo=sysinfo.data.detailInfo
	$rootScope.$broadcast( 'sysinfo.update.link' );
  });
  $scope.pcid = pcid
	$scope.date=[]
	$scope.cpudata=[]
	$scope.memdata=[]
	$scope.detailInfo={}
	$scope.modalContent={}
	$scope.procs_order="pid"
	$scope.procs=[]

	$scope.stopOrRun='run'
	$scope.SORBackground = {background:'#00FF7F'}
	$scope.switchStopRun = function(){
		$scope.stopOrRun = !$rootScope.stopOrRun ? 'run' : 'stop'
		$scope.SORBackground.background =  !$rootScope.stopOrRun ? '#53FF53' : '#FF5151'
		$rootScope.stopOrRun = !$rootScope.stopOrRun
    if($rootScope.stopOrRun) setProcInte()
    else colseProcInte()
	}
  $scope.showModal=function(proc){
    $scope.modalContent=proc
	  var pid = proc.pid
	  $http({
		  url:'/'+pcid+'/proc/'+pid+'/other',
		  method:'get',
	  }).success(function(data){
		  $scope.modalContent.port = data.port
		  $scope.modalContent.io = data.io
		  $scope.$apply()
	  })
    $('#myModal').modal('show')
  }

  var getProcs = function(){
    $http({
			url:'/'+pcid+'/info/proc',
			method:'get',
		}).success(function(data){
			$scope.procs=data
		})
  }

  var setProcInte = function(){
    $scope.procIntervalID = setInterval(getProcs,60000)
  }
  var colseProcInte = function(){
    clearInterval($scope.procIntervalID)
  }
  getProcs()
  setProcInte()

}]);
