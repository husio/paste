(function () {
  window.addEventListener('load', function () {
      var t = document.getElementById('content');
      var resizeContentInput = function () {
        t.style.height = window.innerHeight - 60 + 'px';
      };

      window.addEventListener('resize', resizeContentInput);
      resizeContentInput();
  }, false);
}());
