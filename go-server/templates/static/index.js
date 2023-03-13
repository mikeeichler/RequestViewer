function jsonToTable(j) {
    let table = "<table><tr><th>Header</th><th>Value</th></tr>";
  for (const k in j) {
    table += `<tr><td>${k}</td><td>${j[k]}</td></tr>`;
    if (k == "timestamp") {
      timestamp = j[k];
    }
  }
  table += '</table>'
  return table;
}

function load(device) {
      const videoObject = new Map();
      var video = document.createElement('video');
      
      //document.write(Objects.keys(video));
      
      //if (video.canPlayType('video/mp4; codecs="avc1.42E01E, mp4a.40.2"')) {
        //document.write('H.264 is supported')
      //}
      if (video.canPlayType('video/mp4; codecs="ap4h.2, mp4a.40.2"')) {
          document.write('<br>ProRes 4:2:2 is supported')
          videoObject.set('PR422', true)
      } else {
        videoObject.set('PR422', false)
        document.write('<br>ProRes 4:2:2 is not supported')
      }
      
      if (video.canPlayType('video/mp4; codecs="ap4h.2, mp4a.40.2"')) {
          videoObject.set('PR4444', true)
          document.write('<br>ProRes 4:4:4:4 is supported')
      } else {
          videoObject.set('PR4444', false)
        document.write('<br>ProRes 4:4:4:4 is not supported')
      }
          var generateWebGLData = function generateWebGLData(gl) {
              var t0 = performance.now();
      
              try {
                  var vShaderTemplate = 'attribute vec2 attrVertex;varying vec2 varyinTexCoordinate;uniform vec2 uniformOffset;void main(){varyinTexCoordinate=attrVertex+uniformOffset;gl_Position=vec4(attrVertex,0,1);}';
                  var fShaderTemplate = 'precision mediump float;varying vec2 varyinTexCoordinate;void main() {gl_FragColor=vec4(varyinTexCoordinate,0,1);}';
                  var vertexPosBuffer = gl.createBuffer();
                  gl.bindBuffer(gl.ARRAY_BUFFER, vertexPosBuffer);
                  var vertices = new Float32Array([-.2, -.9, 0, .4, -.26, 0, 0, .732134444, 0]);
                  gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);
                  vertexPosBuffer.itemSize = 3;
                  vertexPosBuffer.numItems = 3;
                  var program = gl.createProgram();
                  var vshader = gl.createShader(gl.VERTEX_SHADER);
                  gl.shaderSource(vshader, vShaderTemplate);
                  gl.compileShader(vshader);
                  var fshader = gl.createShader(gl.FRAGMENT_SHADER);
                  gl.shaderSource(fshader, fShaderTemplate);
                  gl.compileShader(fshader);
                  gl.attachShader(program, vshader);
                  gl.attachShader(program, fshader);
                  gl.linkProgram(program);
                  gl.useProgram(program);
                  program.vertexPosAttrib = gl.getAttribLocation(program, 'attrVertex');
                  program.offsetUniform = gl.getUniformLocation(program, 'uniformOffset');
                  gl.enableVertexAttribArray(program.vertexPosArray);
                  gl.vertexAttribPointer(program.vertexPosAttrib, vertexPosBuffer.itemSize, gl.FLOAT, !1, 0, 0);
                  gl.uniform2f(program.offsetUniform, 1, 1);
                  gl.drawArrays(gl.TRIANGLE_STRIP, 0, vertexPosBuffer.numItems);
      
                  if (gl.canvas != null) {
                      var t1 = performance.now();
      
                      if (window.amiunique_degub) {
                          console.log("execution time for webgl data = ".concat(t1 - t0, " ms"));
                      }
      
                      return gl.canvas.toDataURL();
                  }
              } catch (e) {
                  var _t3 = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl data = ".concat(_t3 - t0, " ms"));
                  }
      
                  return 'Not supported';
              }
          };
      
          var fa2s = function fa2s(fa) {
              gl.clearColor(0.0, 0.0, 0.0, 1.0);
              gl.enable(gl.DEPTH_TEST);
              gl.depthFunc(gl.LEQUAL);
              gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
              return '[' + fa[0] + ', ' + fa[1] + ']';
          };
      
          var getGeneralParameters = function getGeneralParameters(webGLParameters, gl) {
              var t0 = performance.now();
              var fa2sParameters = ['ALIASED_LINE_WIDTH_RANGE', 'ALIASED_POINT_SIZE_RANGE', 'MAX_VIEWPORT_DIMS'];
              var generalParameterNames = ['ALPHA_BITS', 'BLUE_BITS', 'DEPTH_BITS', 'GREEN_BITS', 'MAX_COMBINED_TEXTURE_IMAGE_UNITS', 'MAX_CUBE_MAP_TEXTURE_SIZE', 'MAX_FRAGMENT_UNIFORM_VECTORS', 'MAX_RENDERBUFFER_SIZE', 'MAX_RENDERBUFFER_SIZE', 'MAX_TEXTURE_IMAGE_UNITS', 'MAX_TEXTURE_SIZE', 'MAX_VARYING_VECTORS', 'MAX_VERTEX_ATTRIBS', 'MAX_VERTEX_TEXTURE_IMAGE_UNITS', 'MAX_VERTEX_UNIFORM_VECTORS', 'RED_BITS', 'RENDERER', 'SHADING_LANGUAGE_VERSION', 'STENCIL_BITS', 'VENDOR', 'VERSION'];
              generalParameters = {};
      
              try {
                  generalParameters['MAX_ANISOTROPY'] = maxAnisotropy(gl);
                  generalParameters['ANTIALIAS'] = gl.getContextAttributes().antialias ? 'yes' : 'no';
                  fa2sParameters.forEach(function (fa2sParameter) {
                      generalParameters[fa2sParameter] = fa2s(gl.getParameter(gl[fa2sParameter]));
                  });
                  generalParameterNames.forEach(function (generalParameterName) {
                      generalParameters[generalParameterName] = gl.getParameter(gl[generalParameterName]);
                  });
                  var t1 = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl general parameters = ".concat(t1 - t0, " ms"));
                  }
      
                  return generalParameters;
              } catch (e) {
                  var _t = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl general parameters = ".concat(_t - t0, " ms"));
                  }
      
                  return window.amiunique_not_supported;
              }
          };
      
          var getShaderPrecisionParameters = function getShaderPrecisionParameters(webGLParameters, gl) {
              var t0 = performance.now();
              var shadersTypes = ['VERTEX_SHADER', 'FRAGMENT_SHADER'];
              var numberTypes = ['HIGH_FLOAT', 'MEDIUM_FLOAT', 'LOW_FLOAT', 'HIGH_INT', 'MEDIUM_INT', 'LOW_INT'];
              var parameters = ['precision', 'rangeMin', 'rangeMax'];
              shadersPrecisionParameters = {};
      
              try {
                  shadersTypes.forEach(function (shaderType) {
                      numberTypes.forEach(function (numberType) {
                          parameters.forEach(function (parameter) {
                              var fullName = shaderType + ' ' + numberType + ' ' + parameter;
                              shadersPrecisionParameters[fullName] = gl.getShaderPrecisionFormat(gl[shaderType], gl[numberType])[parameter];
                          });
                      });
                  });
                  var t1 = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl shaders precision parameters = ".concat(t1 - t0, " ms"));
                  }
      
                  return shadersPrecisionParameters;
              } catch (e) {
                  var _t2 = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl shaders precision parameters = ".concat(_t2 - t0, " ms"));
                  }
      
                  return window.amiunique_not_supported;
              }
          };
      
          var maxAnisotropy = function maxAnisotropy(gl) {
              var anisotropy;
              var ext = gl.getExtension('EXT_texture_filter_anisotropic') || gl.getExtension('WEBKIT_EXT_texture_filter_anisotropic') || gl.getExtension('MOZ_EXT_texture_filter_anisotropic');
              return ext ? (anisotropy = gl.getParameter(ext.MAX_TEXTURE_MAX_ANISOTROPY_EXT), 0 === anisotropy && (anisotropy = 2), anisotropy) : null;
          };
      
          try {
              canvas = document.createElement('canvas');
              var gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl');
      
              if (gl.getSupportedExtensions().indexOf('WEBGL_debug_renderer_info') >= 0) {
                  try {
                      webGLVendor = gl.getParameter(gl.getExtension('WEBGL_debug_renderer_info').UNMASKED_VENDOR_WEBGL);
                  } catch (e) {
                      webGLVendor = window.amiunique_not_supported;
                  }
      
                  try {
                      webGLRenderer = gl.getParameter(gl.getExtension('WEBGL_debug_renderer_info').UNMASKED_RENDERER_WEBGL);
                  } catch (e) {
                      webGLRenderer = window.amiunique_not_supported;
                  }
      
                  var t1 = performance.now();
      
                  if (window.amiunique_degub) {
                      console.log("execution time for webgl vendor and webgl renderer = ".concat(t1 - t0, " ms"));
                  }
      
                  webGLData = generateWebGLData(gl);
      
                  webGLDataParameters = {}
                  webGLDataParameters["hash"] = CryptoJS.MD5(webGLData).toString()
                  webGLDataParameters["OriginalData"] = webGLData
      
                  document.write("<pre id=\"json\">" + JSON.stringify(webGLDataParameters, null, 4) + "</pre>");
      
                  document.write("Image: ");
                  var image = new Image();
                  image.src = webGLData;
                  document.body.appendChild(image);
      
                  webGLParameters = {};
                  webGLParameters['extensions'] = gl.getSupportedExtensions();
                  webGLParameters['general'] = getGeneralParameters(webGLParameters, gl);
                  webGLParameters['shaderPrecision'] = getShaderPrecisionParameters(webGLParameters, gl);
      
                  document.write("<pre id=\"json\">" + JSON.stringify(webGLParameters, null, 4) + "</pre>");
              } else {gg
                  webGLVendor = window.amiunique_not_supported;
                  webGLRenderer = window.amiunique_not_supported;
                  webGLParameters = window.amiunique_not_supported;
                  webGLData = window.amiunique_not_supported;
              }
          } catch (e) {
              webGLVendor = window.amiunique_not_supported;
              webGLRenderer = window.amiunique_not_supported;
              webGLParameters = window.amiunique_not_supported;
              webGLData = window.amiunique_not_supported;
          }
    videoObject.set('webglHash', webGLDataParameters.hash);
    fetch("/api", {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'X-PR422': videoObject.get('PR422'),
            'X-PR4444': videoObject.get('PR4444'),
            'X-WEBGL-HASH': videoObject.get('webglHash'),
            'X-ACTUAL-DEVICE': device
        }},)
    .then((response) => response.json())
    .then((data) => {
        // console.log(data)
        // document.getElementById("content").innerHTML = jsonToTable(data);
        document.write(jsonToTable(data))
    })
};