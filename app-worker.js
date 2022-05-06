const cacheName = "app-" + "334bccbe67478565006488dc5a8e568f3a0a36d6";

self.addEventListener("install", event => {
  console.log("installing app worker 334bccbe67478565006488dc5a8e568f3a0a36d6");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/budoux-go",
          "/budoux-go/app.css",
          "/budoux-go/app.js",
          "/budoux-go/manifest.webmanifest",
          "/budoux-go/wasm_exec.js",
          "/budoux-go/web/app.wasm",
          "https://github.com/identicons/sg0hsmt.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 334bccbe67478565006488dc5a8e568f3a0a36d6 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
