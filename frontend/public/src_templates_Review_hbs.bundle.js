/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(self["webpackChunk"] = self["webpackChunk"] || []).push([["src_templates_Review_hbs"],{

/***/ "./src/templates/Review.hbs":
/*!**********************************!*\
  !*** ./src/templates/Review.hbs ***!
  \**********************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

eval("var Handlebars = __webpack_require__(/*! ../../node_modules/handlebars/runtime.js */ \"./node_modules/handlebars/runtime.js\");\nfunction __default(obj) { return obj && (obj.__esModule ? obj[\"default\"] : obj); }\nmodule.exports = (Handlebars[\"default\"] || Handlebars).template({\"1\":function(container,depth0,helpers,partials,data) {\n    return \"               <button class=\\\"button-edit-review\\\">Изменить отзыв</button>\\n               <button class=\\\"button-delete-review button-danger\\\">Удалить отзыв</button>\\n\";\n},\"compiler\":[8,\">= 4.3.0\"],\"main\":function(container,depth0,helpers,partials,data) {\n    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {\n        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {\n          return parent[propertyName];\n        }\n        return undefined\n    };\n\n  return \"<div class=\\\"review\\\" id=\\\"review-\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"reviewContent\") : depth0)) != null ? lookupProperty(stack1,\"id\") : stack1), depth0))\n    + \"\\\">\\n     <div class=\\\"review-top\\\">\\n          <img src=\\\"\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"reviewContent\") : depth0)) != null ? lookupProperty(stack1,\"avatar\") : stack1), depth0))\n    + \"\\\">\\n          <p class=\\\"review-username\\\">\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"reviewContent\") : depth0)) != null ? lookupProperty(stack1,\"username\") : stack1), depth0))\n    + \"</p>\\n     </div>\\n     <div class=\\\"review-content\\\">\\n          <p class=\\\"review-text\\\">\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"reviewContent\") : depth0)) != null ? lookupProperty(stack1,\"feedback\") : stack1), depth0))\n    + \"</p>\\n     </div>\\n     <div class=\\\"rating\\\">\\n\\n     </div>\\n     <div class=\\\"review-buttons\\\">\\n\"\n    + ((stack1 = lookupProperty(helpers,\"if\").call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,\"isOwn\") : depth0),{\"name\":\"if\",\"hash\":{},\"fn\":container.program(1, data, 0),\"inverse\":container.noop,\"data\":data,\"loc\":{\"start\":{\"line\":13,\"column\":10},\"end\":{\"line\":16,\"column\":17}}})) != null ? stack1 : \"\")\n    + \"     </div>\\n</div>\";\n},\"useData\":true});\n\n//# sourceURL=webpack:///./src/templates/Review.hbs?");

/***/ })

}]);