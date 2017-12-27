webpackJsonp([0],{

/***/ 496:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_13_6_1_vue_loader_lib_selector_type_script_index_0_bustCache_node_modules_1_0_0_iview_loader_index_js_ref_0_1_index_vue__ = __webpack_require__(497);
/* empty harmony namespace reexport */
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__node_modules_13_6_1_vue_loader_lib_template_compiler_index_id_data_v_436ef8ce_hasScoped_true_buble_transforms_node_modules_13_6_1_vue_loader_lib_selector_type_template_index_0_bustCache_node_modules_1_0_0_iview_loader_index_js_ref_0_1_index_vue__ = __webpack_require__(499);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(498)
}
var normalizeComponent = __webpack_require__(178)
/* script */


/* template */

/* template functional */
var __vue_template_functional__ = false
/* styles */
var __vue_styles__ = injectStyle
/* scopeId */
var __vue_scopeId__ = "data-v-436ef8ce"
/* moduleIdentifier (server only) */
var __vue_module_identifier__ = null
var Component = normalizeComponent(
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_13_6_1_vue_loader_lib_selector_type_script_index_0_bustCache_node_modules_1_0_0_iview_loader_index_js_ref_0_1_index_vue__["a" /* default */],
  __WEBPACK_IMPORTED_MODULE_1__node_modules_13_6_1_vue_loader_lib_template_compiler_index_id_data_v_436ef8ce_hasScoped_true_buble_transforms_node_modules_13_6_1_vue_loader_lib_selector_type_template_index_0_bustCache_node_modules_1_0_0_iview_loader_index_js_ref_0_1_index_vue__["a" /* default */],
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "src\\views\\index.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-436ef8ce", Component.options)
  } else {
    hotAPI.reload("data-v-436ef8ce", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 497:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

/* harmony default export */ __webpack_exports__["a"] = ({
    methods: {
        handleStart: function handleStart() {
            this.$Modal.info({
                title: 'Bravo',
                content: 'Now, enjoy the convenience of iView.'
            });
        }
    }
});

/***/ }),

/***/ 498:
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),

/***/ 499:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
var render = function() {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c(
    "div",
    { staticClass: "index" },
    [
      _c(
        "Row",
        { attrs: { type: "flex", justify: "center", align: "middle" } },
        [
          _c("Col", { attrs: { span: "24" } }, [
            _c("h1", [
              _c("img", {
                attrs: {
                  src:
                    "https://raw.githubusercontent.com/iview/iview/master/assets/logo.png"
                }
              })
            ]),
            _vm._v(" "),
            _c(
              "h2",
              [
                _c("p", [_vm._v("Welcome to your iView app!")]),
                _vm._v(" "),
                _c(
                  "Button",
                  { attrs: { type: "ghost" }, on: { click: _vm.handleStart } },
                  [_vm._v("Start iView")]
                )
              ],
              1
            )
          ])
        ],
        1
      )
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true
var esExports = { render: render, staticRenderFns: staticRenderFns }
/* harmony default export */ __webpack_exports__["a"] = (esExports);
if (false) {
  module.hot.accept()
  if (module.hot.data) {
    require("vue-loader/node_modules/vue-hot-reload-api")      .rerender("data-v-436ef8ce", esExports)
  }
}

/***/ })

});
//# sourceMappingURL=0.chunk.js.map