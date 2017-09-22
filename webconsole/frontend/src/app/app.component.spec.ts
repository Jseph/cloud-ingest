import { TestBed, async } from '@angular/core/testing';
import { ActivatedRoute, Params, NavigationExtras } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { Observable } from 'rxjs/Observable';
import { AppComponent } from './app.component';
import { AngularMaterialImporterModule } from './angular-material-importer.module';
import { AuthService } from './auth.service';
import { UserProfile } from './auth.resources';
import { NoopAnimationsModule} from '@angular/platform-browser/animations';
import { MdSnackBar } from '@angular/material';
import 'rxjs/add/observable/of';

const FAKE_USER = 'Fake User';
const FAKE_AUTH = 'Fake Auth';

class MockAuthService extends AuthService {
  isSignedIn = true;

  fakeUser: UserProfile = {
    Name: FAKE_USER
  };

  init() { }

  loadSignInStatus(): Promise<boolean> {
    return Promise.resolve(this.isSignedIn);
  }

  getAuthorizationHeader(): string {
    return FAKE_AUTH;
  }

  getCurrentUser(): UserProfile  {
    return this.fakeUser;
  }
}

class MdSnackBarStub {
  open = jasmine.createSpy('open');
}

let activatedRouteStub: ActivatedRoute;
let mdSnackBarStub: MdSnackBarStub;

describe('AppComponent', () => {
  const mockAuthService = new MockAuthService();

  beforeEach(async(() => {
    mockAuthService.isSignedIn = true;
    activatedRouteStub = new ActivatedRoute();
    activatedRouteStub.queryParams = Observable.of({project: 'fakeProjectId'});
    mdSnackBarStub = new MdSnackBarStub();

    TestBed.configureTestingModule({
      declarations: [
        AppComponent
      ],
      providers: [
        {provide: AuthService, useValue: mockAuthService},
        {provide: ActivatedRoute, useValue: activatedRouteStub},
        {provide: MdSnackBar, useValue: mdSnackBarStub},
      ],
      imports: [
        RouterTestingModule,
        NoopAnimationsModule,
        FormsModule,
        AngularMaterialImporterModule
      ],
    }).compileComponents();
  }));

  it('should create the app', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));

  it('should render title in a h1 tag', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      expect(compiled.querySelector('h1').textContent).
          toContain(`Ingest Web Console - ${FAKE_USER}`);
    });
  }));

  it('should contain two links and signout', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      expect(compiled.querySelectorAll('a').length).toBe(2);

      const signOutButton = compiled.querySelector('button');
      expect(signOutButton).not.toBeNull();
      expect(signOutButton.textContent).toContain('Signout');
    });
  }));

  it('should contain a Job Configs link', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('#jobconfigslink');
      expect(element).not.toBeNull();
      expect(element.textContent).toContain('Job Configs');
      expect(element.getAttribute('queryParamsHandling')).toBe('merge');
    });
  }));

  it('should contain an infrastructure link', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('#infrastructure-link');
      expect(element).not.toBeNull();
      expect(element.textContent).toContain('Infrastructure');
      expect(element.getAttribute('queryParamsHandling')).toBe('merge');
    });
  }));

  it('should not show links, and show sign in button if not signed in',
     async(() => {
    mockAuthService.isSignedIn = false;
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      expect(compiled.querySelectorAll('a').length).toBe(0);
      const signInButton = compiled.querySelector('button');
      expect(signInButton).not.toBeNull();
      expect(signInButton.textContent).toContain('Sign In');
    });
  }));

  it('should contain a toolbar', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('md-toolbar');
      expect(element).not.toBeNull();
    });
  }));

  it('should contain a side nav', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('md-sidenav');
      expect(element).not.toBeNull();
    });
  }));

  it('should navigate to the job configs page with project id', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    app.gcsProjectId = 'fakeGcsProjectId';
    const navigateSpy = spyOn((<any>app).router, 'navigate');
    const fakeNavigationExtras: NavigationExtras = {
      queryParams: { project: 'fakeGcsProjectId' }
    };

    app.onProjectSelectSubmit();
    expect(navigateSpy).toHaveBeenCalledWith(['/jobconfigs'], fakeNavigationExtras);
  }));

  it('should show a project selection input if no project is selected', async(() => {
    activatedRouteStub.queryParams = Observable.of({project: ''});
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const projectInput = compiled.querySelector('#projectId');
      expect(projectInput).not.toBeNull();
    });
  }));

  it('should open a snackbar on sign in failure', async(() => {
    spyOn(mockAuthService, 'signIn').and.returnValue(Promise.reject({error: 'fakeSignInFailedMessage'}));
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    app.signIn();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      expect(mdSnackBarStub.open).toHaveBeenCalled();
      expect(mdSnackBarStub.open.calls.first().args[0]).toMatch('fakeSignInFailedMessage');
    });
  }));

  it('should open a snackbar on sign out failure', async(() => {
    spyOn(mockAuthService, 'signOut').and.returnValue(Promise.reject({error: 'fakeSignOutFailedMessage'}));
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    app.signOut();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      expect(mdSnackBarStub.open).toHaveBeenCalled();
      expect(mdSnackBarStub.open.calls.first().args[0]).toMatch('fakeSignOutFailedMessage');
    });
  }));

  it('should open a snackbar on loadSignInStatus failure when initiating the component', async(() => {
    spyOn(mockAuthService, 'loadSignInStatus').and.returnValue(Promise.reject({error: 'fakeLoadSignInFailMessage'}));
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      expect(mdSnackBarStub.open).toHaveBeenCalled();
      expect(mdSnackBarStub.open.calls.first().args[0]).toMatch('fakeLoadSignInFailMessage');
    });
  }));
});
