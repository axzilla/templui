// https://cdn.jsdelivr.net/npm/@floating-ui/dom@1.7.0
!(function (t, e) {
  "object" == typeof exports && "undefined" != typeof module
    ? e(exports, require("./floating_ui_core"))
    : "function" == typeof define && define.amd
    ? define(["exports", "./floatingUICore"], e)
    : e(
        ((t =
          "undefined" != typeof globalThis
            ? globalThis
            : t || self).FloatingUIDOM = {}),
        t.FloatingUICore
      );
})(this, function (t, e) {
  "use strict";
  const n = Math.min,
    o = Math.max,
    i = Math.round,
    r = Math.floor,
    c = (t) => ({ x: t, y: t });
  function l() {
    return "undefined" != typeof window;
  }
  function s(t) {
    return a(t) ? (t.nodeName || "").toLowerCase() : "#document";
  }
  function f(t) {
    var e;
    return (
      (null == t || null == (e = t.ownerDocument) ? void 0 : e.defaultView) ||
      window
    );
  }
  function u(t) {
    var e;
    return null ==
      (e = (a(t) ? t.ownerDocument : t.document) || window.document)
      ? void 0
      : e.documentElement;
  }
  function a(t) {
    return !!l() && (t instanceof Node || t instanceof f(t).Node);
  }
  function d(t) {
    return !!l() && (t instanceof Element || t instanceof f(t).Element);
  }
  function h(t) {
    return !!l() && (t instanceof HTMLElement || t instanceof f(t).HTMLElement);
  }
  function p(t) {
    return (
      !(!l() || "undefined" == typeof ShadowRoot) &&
      (t instanceof ShadowRoot || t instanceof f(t).ShadowRoot)
    );
  }
  function g(t) {
    const { overflow: e, overflowX: n, overflowY: o, display: i } = b(t);
    return (
      /auto|scroll|overlay|hidden|clip/.test(e + o + n) &&
      !["inline", "contents"].includes(i)
    );
  }
  function m(t) {
    return ["table", "td", "th"].includes(s(t));
  }
  function y(t) {
    return [":popover-open", ":modal"].some((e) => {
      try {
        return t.matches(e);
      } catch (t) {
        return !1;
      }
    });
  }
  function w(t) {
    const e = x(),
      n = d(t) ? b(t) : t;
    return (
      ["transform", "translate", "scale", "rotate", "perspective"].some(
        (t) => !!n[t] && "none" !== n[t]
      ) ||
      (!!n.containerType && "normal" !== n.containerType) ||
      (!e && !!n.backdropFilter && "none" !== n.backdropFilter) ||
      (!e && !!n.filter && "none" !== n.filter) ||
      [
        "transform",
        "translate",
        "scale",
        "rotate",
        "perspective",
        "filter",
      ].some((t) => (n.willChange || "").includes(t)) ||
      ["paint", "layout", "strict", "content"].some((t) =>
        (n.contain || "").includes(t)
      )
    );
  }
  function x() {
    return (
      !("undefined" == typeof CSS || !CSS.supports) &&
      CSS.supports("-webkit-backdrop-filter", "none")
    );
  }
  function v(t) {
    return ["html", "body", "#document"].includes(s(t));
  }
  function b(t) {
    return f(t).getComputedStyle(t);
  }
  function T(t) {
    return d(t)
      ? { scrollLeft: t.scrollLeft, scrollTop: t.scrollTop }
      : { scrollLeft: t.scrollX, scrollTop: t.scrollY };
  }
  function L(t) {
    if ("html" === s(t)) return t;
    const e = t.assignedSlot || t.parentNode || (p(t) && t.host) || u(t);
    return p(e) ? e.host : e;
  }
  function R(t) {
    const e = L(t);
    return v(e)
      ? t.ownerDocument
        ? t.ownerDocument.body
        : t.body
      : h(e) && g(e)
      ? e
      : R(e);
  }
  function C(t, e, n) {
    var o;
    void 0 === e && (e = []), void 0 === n && (n = !0);
    const i = R(t),
      r = i === (null == (o = t.ownerDocument) ? void 0 : o.body),
      c = f(i);
    if (r) {
      const t = E(c);
      return e.concat(
        c,
        c.visualViewport || [],
        g(i) ? i : [],
        t && n ? C(t) : []
      );
    }
    return e.concat(i, C(i, [], n));
  }
  function E(t) {
    return t.parent && Object.getPrototypeOf(t.parent) ? t.frameElement : null;
  }
  function S(t) {
    const e = b(t);
    let n = parseFloat(e.width) || 0,
      o = parseFloat(e.height) || 0;
    const r = h(t),
      c = r ? t.offsetWidth : n,
      l = r ? t.offsetHeight : o,
      s = i(n) !== c || i(o) !== l;
    return s && ((n = c), (o = l)), { width: n, height: o, $: s };
  }
  function F(t) {
    return d(t) ? t : t.contextElement;
  }
  function O(t) {
    const e = F(t);
    if (!h(e)) return c(1);
    const n = e.getBoundingClientRect(),
      { width: o, height: r, $: l } = S(e);
    let s = (l ? i(n.width) : n.width) / o,
      f = (l ? i(n.height) : n.height) / r;
    return (
      (s && Number.isFinite(s)) || (s = 1),
      (f && Number.isFinite(f)) || (f = 1),
      { x: s, y: f }
    );
  }
  const D = c(0);
  function H(t) {
    const e = f(t);
    return x() && e.visualViewport
      ? { x: e.visualViewport.offsetLeft, y: e.visualViewport.offsetTop }
      : D;
  }
  function P(t, n, o, i) {
    void 0 === n && (n = !1), void 0 === o && (o = !1);
    const r = t.getBoundingClientRect(),
      l = F(t);
    let s = c(1);
    n && (i ? d(i) && (s = O(i)) : (s = O(t)));
    const u = (function (t, e, n) {
      return void 0 === e && (e = !1), !(!n || (e && n !== f(t))) && e;
    })(l, o, i)
      ? H(l)
      : c(0);
    let a = (r.left + u.x) / s.x,
      h = (r.top + u.y) / s.y,
      p = r.width / s.x,
      g = r.height / s.y;
    if (l) {
      const t = f(l),
        e = i && d(i) ? f(i) : i;
      let n = t,
        o = E(n);
      for (; o && i && e !== n; ) {
        const t = O(o),
          e = o.getBoundingClientRect(),
          i = b(o),
          r = e.left + (o.clientLeft + parseFloat(i.paddingLeft)) * t.x,
          c = e.top + (o.clientTop + parseFloat(i.paddingTop)) * t.y;
        (a *= t.x),
          (h *= t.y),
          (p *= t.x),
          (g *= t.y),
          (a += r),
          (h += c),
          (n = f(o)),
          (o = E(n));
      }
    }
    return e.rectToClientRect({ width: p, height: g, x: a, y: h });
  }
  function W(t, e) {
    const n = T(t).scrollLeft;
    return e ? e.left + n : P(u(t)).left + n;
  }
  function M(t, e, n) {
    void 0 === n && (n = !1);
    const o = t.getBoundingClientRect();
    return {
      x: o.left + e.scrollLeft - (n ? 0 : W(t, o)),
      y: o.top + e.scrollTop,
    };
  }
  function z(t, n, i) {
    let r;
    if ("viewport" === n)
      r = (function (t, e) {
        const n = f(t),
          o = u(t),
          i = n.visualViewport;
        let r = o.clientWidth,
          c = o.clientHeight,
          l = 0,
          s = 0;
        if (i) {
          (r = i.width), (c = i.height);
          const t = x();
          (!t || (t && "fixed" === e)) &&
            ((l = i.offsetLeft), (s = i.offsetTop));
        }
        return { width: r, height: c, x: l, y: s };
      })(t, i);
    else if ("document" === n)
      r = (function (t) {
        const e = u(t),
          n = T(t),
          i = t.ownerDocument.body,
          r = o(e.scrollWidth, e.clientWidth, i.scrollWidth, i.clientWidth),
          c = o(e.scrollHeight, e.clientHeight, i.scrollHeight, i.clientHeight);
        let l = -n.scrollLeft + W(t);
        const s = -n.scrollTop;
        return (
          "rtl" === b(i).direction &&
            (l += o(e.clientWidth, i.clientWidth) - r),
          { width: r, height: c, x: l, y: s }
        );
      })(u(t));
    else if (d(n))
      r = (function (t, e) {
        const n = P(t, !0, "fixed" === e),
          o = n.top + t.clientTop,
          i = n.left + t.clientLeft,
          r = h(t) ? O(t) : c(1);
        return {
          width: t.clientWidth * r.x,
          height: t.clientHeight * r.y,
          x: i * r.x,
          y: o * r.y,
        };
      })(n, i);
    else {
      const e = H(t);
      r = { x: n.x - e.x, y: n.y - e.y, width: n.width, height: n.height };
    }
    return e.rectToClientRect(r);
  }
  function A(t, e) {
    const n = L(t);
    return (
      !(n === e || !d(n) || v(n)) && ("fixed" === b(n).position || A(n, e))
    );
  }
  function B(t, e, n) {
    const o = h(e),
      i = u(e),
      r = "fixed" === n,
      l = P(t, !0, r, e);
    let f = { scrollLeft: 0, scrollTop: 0 };
    const a = c(0);
    function d() {
      a.x = W(i);
    }
    if (o || (!o && !r))
      if ((("body" !== s(e) || g(i)) && (f = T(e)), o)) {
        const t = P(e, !0, r, e);
        (a.x = t.x + e.clientLeft), (a.y = t.y + e.clientTop);
      } else i && d();
    r && !o && i && d();
    const p = !i || o || r ? c(0) : M(i, f);
    return {
      x: l.left + f.scrollLeft - a.x - p.x,
      y: l.top + f.scrollTop - a.y - p.y,
      width: l.width,
      height: l.height,
    };
  }
  function V(t) {
    return "static" === b(t).position;
  }
  function N(t, e) {
    if (!h(t) || "fixed" === b(t).position) return null;
    if (e) return e(t);
    let n = t.offsetParent;
    return u(t) === n && (n = n.ownerDocument.body), n;
  }
  function I(t, e) {
    const n = f(t);
    if (y(t)) return n;
    if (!h(t)) {
      let e = L(t);
      for (; e && !v(e); ) {
        if (d(e) && !V(e)) return e;
        e = L(e);
      }
      return n;
    }
    let o = N(t, e);
    for (; o && m(o) && V(o); ) o = N(o, e);
    return o && v(o) && V(o) && !w(o)
      ? n
      : o ||
          (function (t) {
            let e = L(t);
            for (; h(e) && !v(e); ) {
              if (w(e)) return e;
              if (y(e)) return null;
              e = L(e);
            }
            return null;
          })(t) ||
          n;
  }
  const k = {
    convertOffsetParentRelativeRectToViewportRelativeRect: function (t) {
      let { elements: e, rect: n, offsetParent: o, strategy: i } = t;
      const r = "fixed" === i,
        l = u(o),
        f = !!e && y(e.floating);
      if (o === l || (f && r)) return n;
      let a = { scrollLeft: 0, scrollTop: 0 },
        d = c(1);
      const p = c(0),
        m = h(o);
      if (
        (m || (!m && !r)) &&
        (("body" !== s(o) || g(l)) && (a = T(o)), h(o))
      ) {
        const t = P(o);
        (d = O(o)), (p.x = t.x + o.clientLeft), (p.y = t.y + o.clientTop);
      }
      const w = !l || m || r ? c(0) : M(l, a, !0);
      return {
        width: n.width * d.x,
        height: n.height * d.y,
        x: n.x * d.x - a.scrollLeft * d.x + p.x + w.x,
        y: n.y * d.y - a.scrollTop * d.y + p.y + w.y,
      };
    },
    getDocumentElement: u,
    getClippingRect: function (t) {
      let { element: e, boundary: i, rootBoundary: r, strategy: c } = t;
      const l = [
          ...("clippingAncestors" === i
            ? y(e)
              ? []
              : (function (t, e) {
                  const n = e.get(t);
                  if (n) return n;
                  let o = C(t, [], !1).filter((t) => d(t) && "body" !== s(t)),
                    i = null;
                  const r = "fixed" === b(t).position;
                  let c = r ? L(t) : t;
                  for (; d(c) && !v(c); ) {
                    const e = b(c),
                      n = w(c);
                    n || "fixed" !== e.position || (i = null),
                      (
                        r
                          ? !n && !i
                          : (!n &&
                              "static" === e.position &&
                              i &&
                              ["absolute", "fixed"].includes(i.position)) ||
                            (g(c) && !n && A(t, c))
                      )
                        ? (o = o.filter((t) => t !== c))
                        : (i = e),
                      (c = L(c));
                  }
                  return e.set(t, o), o;
                })(e, this._c)
            : [].concat(i)),
          r,
        ],
        f = l[0],
        u = l.reduce((t, i) => {
          const r = z(e, i, c);
          return (
            (t.top = o(r.top, t.top)),
            (t.right = n(r.right, t.right)),
            (t.bottom = n(r.bottom, t.bottom)),
            (t.left = o(r.left, t.left)),
            t
          );
        }, z(e, f, c));
      return {
        width: u.right - u.left,
        height: u.bottom - u.top,
        x: u.left,
        y: u.top,
      };
    },
    getOffsetParent: I,
    getElementRects: async function (t) {
      const e = this.getOffsetParent || I,
        n = this.getDimensions,
        o = await n(t.floating);
      return {
        reference: B(t.reference, await e(t.floating), t.strategy),
        floating: { x: 0, y: 0, width: o.width, height: o.height },
      };
    },
    getClientRects: function (t) {
      return Array.from(t.getClientRects());
    },
    getDimensions: function (t) {
      const { width: e, height: n } = S(t);
      return { width: e, height: n };
    },
    getScale: O,
    isElement: d,
    isRTL: function (t) {
      return "rtl" === b(t).direction;
    },
  };
  function q(t, e) {
    return (
      t.x === e.x && t.y === e.y && t.width === e.width && t.height === e.height
    );
  }
  const U = e.detectOverflow,
    j = e.offset,
    X = e.autoPlacement,
    Y = e.shift,
    $ = e.flip,
    _ = e.size,
    G = e.hide,
    J = e.arrow,
    K = e.inline,
    Q = e.limitShift;
  (t.arrow = J),
    (t.autoPlacement = X),
    (t.autoUpdate = function (t, e, i, c) {
      void 0 === c && (c = {});
      const {
          ancestorScroll: l = !0,
          ancestorResize: s = !0,
          elementResize: f = "function" == typeof ResizeObserver,
          layoutShift: a = "function" == typeof IntersectionObserver,
          animationFrame: d = !1,
        } = c,
        h = F(t),
        p = l || s ? [...(h ? C(h) : []), ...C(e)] : [];
      p.forEach((t) => {
        l && t.addEventListener("scroll", i, { passive: !0 }),
          s && t.addEventListener("resize", i);
      });
      const g =
        h && a
          ? (function (t, e) {
              let i,
                c = null;
              const l = u(t);
              function s() {
                var t;
                clearTimeout(i), null == (t = c) || t.disconnect(), (c = null);
              }
              return (
                (function f(u, a) {
                  void 0 === u && (u = !1), void 0 === a && (a = 1), s();
                  const d = t.getBoundingClientRect(),
                    { left: h, top: p, width: g, height: m } = d;
                  if ((u || e(), !g || !m)) return;
                  const y = {
                    rootMargin:
                      -r(p) +
                      "px " +
                      -r(l.clientWidth - (h + g)) +
                      "px " +
                      -r(l.clientHeight - (p + m)) +
                      "px " +
                      -r(h) +
                      "px",
                    threshold: o(0, n(1, a)) || 1,
                  };
                  let w = !0;
                  function x(e) {
                    const n = e[0].intersectionRatio;
                    if (n !== a) {
                      if (!w) return f();
                      n
                        ? f(!1, n)
                        : (i = setTimeout(() => {
                            f(!1, 1e-7);
                          }, 1e3));
                    }
                    1 !== n || q(d, t.getBoundingClientRect()) || f(), (w = !1);
                  }
                  try {
                    c = new IntersectionObserver(x, {
                      ...y,
                      root: l.ownerDocument,
                    });
                  } catch (t) {
                    c = new IntersectionObserver(x, y);
                  }
                  c.observe(t);
                })(!0),
                s
              );
            })(h, i)
          : null;
      let m,
        y = -1,
        w = null;
      f &&
        ((w = new ResizeObserver((t) => {
          let [n] = t;
          n &&
            n.target === h &&
            w &&
            (w.unobserve(e),
            cancelAnimationFrame(y),
            (y = requestAnimationFrame(() => {
              var t;
              null == (t = w) || t.observe(e);
            }))),
            i();
        })),
        h && !d && w.observe(h),
        w.observe(e));
      let x = d ? P(t) : null;
      return (
        d &&
          (function e() {
            const n = P(t);
            x && !q(x, n) && i();
            (x = n), (m = requestAnimationFrame(e));
          })(),
        i(),
        () => {
          var t;
          p.forEach((t) => {
            l && t.removeEventListener("scroll", i),
              s && t.removeEventListener("resize", i);
          }),
            null == g || g(),
            null == (t = w) || t.disconnect(),
            (w = null),
            d && cancelAnimationFrame(m);
        }
      );
    }),
    (t.computePosition = (t, n, o) => {
      const i = new Map(),
        r = { platform: k, ...o },
        c = { ...r.platform, _c: i };
      return e.computePosition(t, n, { ...r, platform: c });
    }),
    (t.detectOverflow = U),
    (t.flip = $),
    (t.getOverflowAncestors = C),
    (t.hide = G),
    (t.inline = K),
    (t.limitShift = Q),
    (t.offset = j),
    (t.platform = k),
    (t.shift = Y),
    (t.size = _);

  // We put this manually here because we need to make sure it's available
  // before the popover component is initialized.
  window.FloatingUIDOM = t;
  return t;
});
