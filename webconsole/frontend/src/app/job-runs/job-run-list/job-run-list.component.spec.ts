import { TestBed, async } from '@angular/core/testing';
import { JobsService } from '../../jobs.service';
import { JobRun } from '../../api.resources';
import { JobRunListComponent } from './job-run-list.component';
import { JobStatusPipe } from '../job-status.pipe';
import { AngularMaterialImporterModule } from '../../angular-material-importer.module';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/throw';
import 'rxjs/add/observable/of';

class JobsServiceStub {
  public getJobRuns = jasmine.createSpy('getJobRuns');
}

const FAKE_JOB_RUNS: JobRun[] = [
  {
    JobConfigId: 'fakeJobConfigId1',
    JobRunId: 'fakeJobRunId1',
    JobCreationTime: '1504833274371000000',
    Status: 0
  },
  {
    JobConfigId: 'fakeJobConfigId2',
    JobRunId: 'fakeJobRunId2',
    JobCreationTime: '1504833274371000000',
    Status: 1
  },
  {
    JobConfigId: 'fakeJobConfigId3',
    JobRunId: 'fakeJobRunId3',
    JobCreationTime: '1504833274371000000',
    Status: 2
  },
  {
    JobConfigId: 'fakeJobConfigId4',
    JobRunId: 'fakeJobRunId4',
    JobCreationTime: '1504833274371000000',
    Status: 3
  }
];

let jobsServiceStub: JobsServiceStub;
const FAKE_HTTP_ERROR = {error: 'FakeErrorTitle', message: 'Fake error message.'};

describe('JobRunListComponent', () => {

  beforeEach(async(() => {
    jobsServiceStub = new JobsServiceStub();
    jobsServiceStub.getJobRuns.and.returnValue(Observable.of(FAKE_JOB_RUNS));

    TestBed.configureTestingModule({
      declarations: [
        JobRunListComponent,
        JobStatusPipe
      ],
      providers: [
        {provide: JobsService, useValue: jobsServiceStub},
      ],
      imports: [
        AngularMaterialImporterModule
      ],
    }).compileComponents();
  }));

  it('should create the job runs component', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    const component = fixture.debugElement.componentInstance;
    expect(component).toBeTruthy();
  }));

  it('should initialize the component with the expected display columns', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    const component = fixture.debugElement.componentInstance;
    expect(component.displayedColumns).toContain('runId');
    expect(component.displayedColumns).toContain('configId');
    expect(component.displayedColumns).toContain('creationTime');
    expect(component.displayedColumns).toContain('status');
  }));

  it('should contain an md table', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('md-table');
      expect(element).not.toBeNull();
    });
  }));

  it('should contain three md rows', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const elements = compiled.querySelectorAll('md-row');
      expect(elements.length).toBe(4);
    });
  }));

  it('should contain the job config id and job id from jobs service', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const elements = compiled.querySelectorAll('md-row');
      expect(elements[0].innerText).toContain('fakeJobConfigId1');
      expect(elements[0].innerText).toContain('fakeJobRunId1');

      expect(elements[1].innerText).toContain('fakeJobConfigId2');
      expect(elements[1].innerText).toContain('fakeJobRunId2');

      expect(elements[2].innerText).toContain('fakeJobConfigId3');
      expect(elements[2].innerText).toContain('fakeJobRunId3');

      expect(elements[3].innerText).toContain('fakeJobConfigId4');
      expect(elements[3].innerText).toContain('fakeJobRunId4');
    });
  }));

  it('should show a human readable status in the row', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const elements = compiled.querySelectorAll('md-row');
      expect(elements[0].innerText).toContain('Not Started');
      expect(elements[1].innerText).toContain('In Progress');
      expect(elements[2].innerText).toContain('Failed');
      expect(elements[3].innerText).toContain('Success');
    });
  }));

  it('should contain a human readable date', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const elements = compiled.querySelectorAll('md-row');
      expect(elements[0].innerText).toMatch(/\b\d{1,2}[/]\d{1,2}[/]\d{4}\b/);
      expect(elements[1].innerText).toMatch(/\b\d{1,2}[/]\d{1,2}[/]\d{4}\b/);
      expect(elements[2].innerText).toMatch(/\b\d{1,2}[/]\d{1,2}[/]\d{4}\b/);
      expect(elements[3].innerText).toMatch(/\b\d{1,2}[/]\d{1,2}[/]\d{4}\b/);
    });
  }));

  it('should display an error message div', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    jobsServiceStub.getJobRuns.and.returnValue(Observable.throw(FAKE_HTTP_ERROR));
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('.ingest-error-message');
      expect(element).not.toBeNull();
    });
  }));

  it('should display the error message in the http error response', async(() => {
    const fixture = TestBed.createComponent(JobRunListComponent);
    jobsServiceStub.getJobRuns.and.returnValue(Observable.throw(FAKE_HTTP_ERROR));
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      const element = compiled.querySelector('.ingest-error-message');
      expect(element.textContent).toContain('FakeErrorTitle');
      expect(element.textContent).toContain('Fake error message.');
    });
  }));
});