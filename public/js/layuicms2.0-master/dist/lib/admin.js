/** layuiAdmin.pro-v1.4.0 LPPL License By https://www.layui.com/admin/ */
;layui.define("view", function (e) {
    var l = layui.view;
    C = {
        v: "1.4.0 pro", req: l.req, exit: l.exit, escape: function (e) {
            return String(e || "").replace(/&(?!#?[a-zA-Z0-9]+;)/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/'/g, "&#39;").replace(/"/g, "&quot;")
        }, on: function (e, a) {
            return layui.onevent.call(this, n.MOD_NAME, e, a)
        }, popup: l.popup, popupRight: function (e) {
            return C.popup.index = layer.open(a.extend({
                type: 1,
                id: "LAY_adminPopupR",
                anim: -1,
                title: !1,
                closeBtn: !1,
                offset: "r",
                shade: .1,
                shadeClose: !0,
                skin: "layui-anim layui-anim-rl layui-layer-adminRight",
                area: "300px"
            }, e))
        }
    };


    e("admin", C)
});