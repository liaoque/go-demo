/**

 @Name：全局配置
 @Author：贤心
 @Site：http://www.layui.com/admin/
 @License：LPPL（layui付费产品协议）
    
 */
 
layui.define(['laytpl', 'layer', 'element', 'util'], function(exports){
  exports('setter', {

    
    debug: true //是否开启调试模式。如开启，接口异常时会抛出异常 URL 等信息
    
    ,interceptor: false //是否开启未登入拦截
    
    //自定义请求字段
    ,request: {
      tokenName: false //自动携带 token 的字段名。可设置 false 不携带。
    }
    
    //自定义响应字段
    ,response: {
      statusName: 'code' //数据状态的字段名称
      ,statusCode: {
        ok: 0 //数据状态一切正常的状态码
        ,logout: -1000 //登录状态失效的状态码
      }
      ,msgName: 'mes' //状态信息的字段名称
      ,dataName: 'data' //数据详情的字段名称
    }

    
    //扩展的第三方模块
    ,extend: [
      'echarts', //echarts 核心包
      'echartsTheme' //echarts 主题
    ]

  });
});
