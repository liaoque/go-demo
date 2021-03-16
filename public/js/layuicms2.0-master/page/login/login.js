layui.config({
    base: './js/layuicms2.0-master/dist/' //指定 layuiAdmin 项目路径
    ,version: '1.4.0'
}).use(['index','form','layer','jquery'],function(){
    var form = layui.form,admin = layui.admin
        layer = parent.layer === undefined ? layui.layer : top.layer
        $ = layui.jquery;

    $("#imgCode img").click(function(){
        this.src = '/captcha?w=116&h=36&t=' + Math.random()
    })

    //登录按钮
    form.on("submit(login)",function(obj){

        $(this).data('text', $(this).text()).text("登录中...").attr("disabled","disabled").addClass("layui-disabled");
        //请求登入接口
        admin.req({
            url: './login' //实际使用请改成服务端真实接口
            ,data: obj.field
            ,type: "PUT"
            ,done: function(res){

                //登入成功的提示与跳转
                layer.msg('登入成功', {
                    offset: '15px'
                    ,icon: 1
                    ,time: 1000
                }, function(){
                    location.hash = '/';
                });
            }
        });

        return false;
    })

    //表单输入效果
    $(".loginBody .input-item").click(function(e){
        e.stopPropagation();
        $(this).addClass("layui-input-focus").find(".layui-input").focus();
    })
    $(".loginBody .layui-form-item .layui-input").focus(function(){
        $(this).parent().addClass("layui-input-focus");
    })
    $(".loginBody .layui-form-item .layui-input").blur(function(){
        $(this).parent().removeClass("layui-input-focus");
        if($(this).val() != ''){
            $(this).parent().addClass("layui-input-active");
        }else{
            $(this).parent().removeClass("layui-input-active");
        }
    })
})
