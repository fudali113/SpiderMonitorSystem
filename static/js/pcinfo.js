var system = angular.module('system', []);

system.directive('cpu', function() {
    return {
		scope:{
			id:"@",
			date:"=",
			cpudata:"=",
			memdata:"="
			},
        restrict: 'E',
        template: '<div class="panel panel-info" style="height:300px;width:600px"><div style="height:300px;width:600px"></div></div>',
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
			memdata:[50]
		},
		addData:function(data){
			this.data.date.push(getNowTimeStr())
			this.data.cpudata.push(data.cpu[0])
			this.data.memdata.push(data.mem.usedpercent)
			$rootScope.$broadcast('sysinfo.update');
		}
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
	setInterval(getSysinfo,5000)
	return service
}]);

system.controller('sysinfoshow',['$rootScope','$scope','$http','sysinfo',function($rootScope,$scope,$http,sysinfo){
	$scope.$on( 'sysinfo.update', function( event ) {
    $scope.date = sysinfo.data.date
		$scope.cpudata = sysinfo.data.cpudata
		$scope.memdata = sysinfo.data.memdata
		$rootScope.$broadcast( 'sysinfo.update.link' );
  });
	$scope.date=[]
	$scope.cpudata=[]
	$scope.memdata=[]
  $scope.modalContent={}
  $scope.procs_order="pid"
  $scope.procs=[
    {pid:11,name:"dsfd",threads:7,memper:13,io:54654},
    {pid:6,name:"fgvfdg",threads:7,memper:13,io:54654},
    {pid:54,name:"yksfd",threads:7,memper:13,io:54654},
    {pid:3,name:"ghfd",threads:7,memper:13,io:54654},
    {pid:1,name:"asfd",threads:7,memper:13,io:54654}
  ]

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
    $('#myModal').modal('show')
  }
  var getProcs = function(){
    $http({
			url:'/'+pcid+'/info/procs',
			method:'get',
		}).success(function(data){
			$scope.procs=data
		})
  }

  var setProcInte = function(){
    $scope.procIntervalID = setInterval(getProcs,5000)
  }
  var colseProcInte = function(){
    clearInterval($scope.procIntervalID)
  }
  setProcInte()

}]);
