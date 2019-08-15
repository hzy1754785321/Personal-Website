var passwdReset = function(){
    layui.use(['layer','form'], function(){
    var layer = layui.layer
    ,form = layui.form
    layer.open({
                        type: 1,
                        shift: 7,
                        title: false,
                        shadeClose: true,
                        closeBtn: false,
                        area:["310px","360px"],
                        skin: 'layui-elem-field',
                        content:"<fieldset class='layui-elem-field site-demo-button' style='margin-top: 20px;'>" +
                                        "<legend>修改密码</legend>" +
                                "</fieldset>" +
                                "<form class='layui-form layui-form-pane' style='margin-top: 20px;' action='' '>" +
                                    "<div class='layui-form-item'  style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>旧密码</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerOldPwd' class='layui-input' name='pwd' placeholder='请输入旧密码' maxlength='16' type='text' autocomplete='off'>" +
                                        "</div>" +
                                    "</div>" +
                                    "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>新密码</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerPwd' class='layui-input' name='pwd' placeholder='请输入新密码' maxlength='16' type='text' autocomplete='off'>" +
                                        "</div>" +
                                    "</div>" +
                                    "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>确认密码</label>" +
                                        "<div class='layui-input-inline'>" +  
                                            "<input id='layerPwd2' class='layui-input' name='pwd' placeholder='请确认密码' maxlength='16' type='text' autocomplete='off'>" +
                                        "</div>" +
                                    "</div>" +
                                "</from>" +
                                "<div style='margin-top: 40px;margin-left: 60px;'>" + 
                                    "<button id ='pwdConfirm' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>确定</button>" +
                                    "<button style='margin-left: 60px;' id='pwdCancel' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>取消</button>" +
                                "</div>"
                });
            $('#pwdCancel').click(function (){
                layer.closeAll();
            });
            $('#pwdConfirm').click(function (){
                var newPwd = $('#layerPwd').val()
                ,newPwd2 =  $('#layerPwd2').val()
                ,oldPwd = $('#layerOldPwd').val();
                if(newPwd != newPwd2){
                    layer.msg("两次输入的密码不一致!");
                    return;
                }
                var JsonData = { sessionID: sessionId,pwd: oldPwd};
                    $.ajax({
                                    type: 'post',
                                    url: '/CheckPassword',
                                    data: JsonData,
                                    success: function(res) {
                                        if(res.status){
                                        
                                            JsonData = { sessionID: sessionId,Password: newPwd,securityType: "0"};
                                            $.ajax({
                                                type: 'post',
                                                url: '/ChangeSecurity',
                                                data: JsonData,
                                                success: function(res) {
                                                    if(res.status){
                                                        layer.msg(res.msg);
                                                        setTimeout(function () {
                                                            layer.closeAll();
                                                        },1500);
                                                    }else{
                                                        layer.msg('修改错误');
                                                    }
                                                }
                                            });
                                        }else{
                                            layer.msg('密码错误');
                                        }
                                    }
                    });
            });
            form.render();           
    });
 }

 var emailReset = function(){
    layui.use(['layer','form'], function(){
        var layer = layui.layer
        ,form = layui.form
        layer.open({
                        type: 1,
                        shift: 7,
                        title: false,
                        shadeClose: true,
                        closeBtn: false,
                        area:["410px","360px"],
                        skin: 'layui-elem-field',
                        content:"<fieldset class='layui-elem-field site-demo-button' style='margin-top: 20px;'>" +
                                    "<legend>绑定邮箱</legend>" +
                                "</fieldset>" +
                                "<form class='layui-form layui-form-pane' style='margin-top: 20px;' action='' '>" +
                                    "<div class='layui-form-item'  style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>邮箱</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerEmail' class='layui-input' name='email' placeholder='请输入邮箱' maxlength='16' type='text' autocomplete='off'>" +
                                        "</div>" +
                                    "</div>" +
                                    "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>验证码</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerCode' class='layui-input' name='code' placeholder='请输入验证码' maxlength='4' type='text' autocomplete='off'>" +
                                        "</div>" +
                                        "<button id ='sendMailBtn' type='button' class='layui-btn layui-btn-normal layui-btn-radius layui-btn-sm'>发送验证邮件</button>" +
                                    "</div>" +
                                "</form>"
        });
    });
}