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
        if($("#emailBind").text() == "已绑定"){
            layer.open({
                type: 1,
                shift: 7,
                title: false,
                shadeClose: true,
                closeBtn: false,
                area:["410px","300px"],
                skin: 'layui-elem-field',
                content:"<fieldset class='layui-elem-field site-demo-button' style='margin-top: 20px;'>" +
                            "<legend>解绑邮箱</legend>" +
                        "</fieldset>" +
                        "<form class='layui-form layui-form-pane' style='margin-top: 20px;' action='' '>" +
                            "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                "<label class='layui-form-label'>当前邮箱：</label>" +
                                
                                "<label id='bindEmail' style='width : 190px;'  class='layui-form-label '></label>" +
                          //      "<text id='bindEmail'></text>" +
                            "</div>" +
                            "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                "<label class='layui-form-label'>验证码</label>" +
                                "<div class='layui-input-inline'>" +                                                               
                                    "<input id='layerCode' class='layui-input' name='code' placeholder='请输入验证码' maxlength='6' type='text' autocomplete='off'>" +
                                "</div>" +
                                "<button id ='sendUnbindMailBtn' style='margin-top: 4px' type='button' class='layui-btn layui-btn-normal layui-btn-radius layui-btn-sm'>发送解绑邮件</button>" +
                            "</div>" +
                        "</form>" +
                        "<div style='margin-top: 40px;margin-left: 60px;'>" + 
                            "<button id ='unBindemailConfirm' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>确定</button>" +
                            "<button style='margin-left: 60px;' id='emailCancel' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>取消</button>" +
                        "</div>"
            });
        }else if($("#emailBind").text() == "未绑定"){
            layer.open({
                        type: 1,
                        shift: 7,
                        title: false,
                        shadeClose: true,
                        closeBtn: false,
                        area:["410px","300px"],
                        skin: 'layui-elem-field',
                        content:"<fieldset class='layui-elem-field site-demo-button' style='margin-top: 20px;'>" +
                                    "<legend>绑定邮箱</legend>" +
                                "</fieldset>" +
                                "<form class='layui-form layui-form-pane' style='margin-top: 20px;' action='' '>" +
                                    "<div class='layui-form-item'  style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>邮箱</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerEmail' class='layui-input' name='email' placeholder='请输入邮箱' type='text' autocomplete='off'>" +
                                        "</div>" +
                                    "</div>" +
                                    "<div class='layui-form-item' style='margin-top: 30px;'>" +
                                        "<label class='layui-form-label'>验证码</label>" +
                                        "<div class='layui-input-inline'>" +                                                               
                                            "<input id='layerCode' class='layui-input' name='code' placeholder='请输入验证码' maxlength='6' type='text' autocomplete='off'>" +
                                        "</div>" +
                                        "<button id ='sendMailBtn' style='margin-top: 4px' type='button' class='layui-btn layui-btn-normal layui-btn-radius layui-btn-sm'>发送验证邮件</button>" +
                                    "</div>" +
                                "</form>" +
                                "<div style='margin-top: 40px;margin-left: 60px;'>" + 
                                    "<button id ='emailConfirm' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>确定</button>" +
                                    "<button style='margin-left: 60px;' id='emailCancel' type='button' class='layui-btn layui-btn-normal layui-btn-radius'>取消</button>" +
                                "</div>"
            });
        }
        $('#emailCancel').click(function (){
            layer.closeAll();
        });
        $('#bindEmail').text($("#emailText").text())
        var emailAddress = $("#emailText").text()
        $('#sendUnbindMailBtn').click(function (){
            var JsonData = { sessionID: sessionId,email: emailAddress,type:"unbind"};
            $.ajax({
                        type: 'post',
                        url: '/SendVerifyMail',
                        data: JsonData,
                        success: function(res) {
                            layer.msg(res.msg);      
                        }
            });
        });
        $('#unBindemailConfirm').click(function (){   
            var emailCode = $('#layerCode').val();
            var emailAddress = $('#bindEmail').val();
            if(emailCode == ""){
                layer.msg("请填写验证码");
                return;
            }
            var JsonData = { sessionID: sessionId,email: emailAddress,cdkey: emailCode,type :"unbind"};
                $.ajax({
                    type: 'post',
                    url: '/VerifyMailCode',
                    data: JsonData,
                    success: function(res) {
            //            layer.msg(res.msg);
                        if(res.status){
                            setTimeout(function () {
                                layer.closeAll();
                            },1500);
                            location.reload() 
                        }
                    }
                });
        });

        $('#sendMailBtn').click(function (){
            var emailAddress = $('#layerEmail').val();
            var sReg = /[_a-zA-Z\d\-\.]+@[_a-zA-Z\d\-]+(\.[_a-zA-Z\d\-]+)+$/;
            if (!sReg.test(emailAddress) ){
                layer.msg("邮箱填写错误！");
                return;     
            }
            var JsonData = { sessionID: sessionId,email: emailAddress,type:"bind",type :"bind"};
                    $.ajax({
                                type: 'post',
                                url: '/SendVerifyMail',
                                data: JsonData,
                                success: function(res) {
                                    layer.msg(res.msg);      
                                }
                    });
        });
        $('#emailConfirm').click(function (){    
            var emailAddress = $('#layerEmail').val();
            var emailCode = $('#layerCode').val();
            if(emailCode == ""){
                layer.msg("请填写验证码");
                return;
            }
            var JsonData = { sessionID: sessionId,email: emailAddress,cdkey: emailCode,type:"bind"};
                $.ajax({
                    type: 'post',
                    url: '/VerifyMailCode',
                    data: JsonData,
                    success: function(res) {
     //                   layer.msg(res.msg);
                        if(res.status){
                            setTimeout(function () {
                                layer.closeAll();
                            },1500);
                            location.reload() 
                        }
                    }
                });
        });
    });
}

var iphoneReset = function(){
    layui.use(['layer'], function(){
        var layer = layui.layer;
        layer.msg("发不起手机验证码！！！")
    });
}