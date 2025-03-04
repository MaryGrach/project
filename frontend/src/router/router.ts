import urls from '@router/urls';
import LoginPage from '@pages/LoginPage/LoginPage';
import PlacesPage from '@pages/PlacesPage/PlacesPage';
import SignupPage from '@pages/SignupPage/SignupPage';
import NotFoundPage from '@pages/NotFoundPage/NotFoundPage';
import ProfilePage from '@pages/ProfilePage/ProfilePage';
import SightPage from '@pages/SightPage/SightPage';
import JourneyPage from '@pages/JourneyPage/JourneyPage';
import AlbumPage from '@pages/AlbumPage/AlbumPage';

const routesList = {
  [urls.base]: PlacesPage,
  [urls.signup]: SignupPage,
  [urls.login]: LoginPage,
  [urls.notfound]: NotFoundPage,
  [urls.profile]: ProfilePage,
  [urls.sight]: SightPage,
  [urls.journey]: JourneyPage,
  // [urls.albums]: AlbumPage,
};

interface Page {
  render: () => void;
}

type Routes = Record<string, new (content: HTMLElement, ...args: any[]) => Page>;

class Router {
  routes: Routes;

  params : {};

  constructor(routes: Routes) {
    this.routes = routes;
    window.addEventListener('popstate', () => this.changeLocation());

    document.addEventListener('click', (e) => {
      let href: string | null;
      const target = e.target as HTMLElement;
      if (target.tagName === 'A' || target.tagName === 'BUTTON') {
        e.preventDefault();
        href = target.getAttribute('href');
      } else if (target.tagName === 'BUTTON' || target.tagName === 'IMG') {
        const parentAnchorElement = target.closest('a');
        if (parentAnchorElement !== null && parentAnchorElement !== undefined) {
          e.preventDefault();
          href = target.closest('a').getAttribute('href');
        }
      } else {
        return;
      }
      if (href) this.go(href);
    });

    document.addEventListener('redirect', (e) => {
      this.go(e.detail.path); // будет рефактор для этой более простой истории
    });

  }

  route(path: string, PageConstructor: new (content: HTMLElement, param?: string) => Page) {
    this.routes[path] = PageConstructor;
    return this;
  }

  go(path: string, params?: string) {
    if (!path.startsWith('/')) {
      path = '/' + path;
    }
    if (window.location.pathname !== path) {
      window.history.pushState({
        params,
      }, '', path);
    }
    this.changeLocation();
  }

  goBack() {
    if (document.referrer) {
      this.go(new URL(document.referrer).pathname);
    } else {
      window.history.back();
    }
  }

  clearContent() {
    let content = document.getElementById('content') as HTMLDivElement;
    if (!content) {
      content = document.createElement('div');
      content.id = 'content';
      const root = document.getElementById('root') as HTMLDivElement;
      root.appendChild(content);
    }
    content.innerHTML = '';
    return content;
  }

  changeLocation() {
    let path = window.location.pathname;
    const PageConstructorFromRoutes = this.routes[path.split('/')[1]];
    const content = this.clearContent();
    const paramArray = path.split('/').slice(2);
    this.params = paramArray;

    if (PageConstructorFromRoutes) {
      const page = new PageConstructorFromRoutes(content, this.params);
      page.render();
    } else {
      const PageConstructorNotFound = this.routes['404'];
      if (PageConstructorNotFound) {
        new PageConstructorNotFound(content).render();
      }
    }
  }
}


const router = new Router(routesList);
export {
  router,
};
