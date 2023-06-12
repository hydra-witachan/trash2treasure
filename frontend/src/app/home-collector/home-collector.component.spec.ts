import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeCollectorComponent } from './home-collector.component';

describe('HomeCollectorComponent', () => {
  let component: HomeCollectorComponent;
  let fixture: ComponentFixture<HomeCollectorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [HomeCollectorComponent]
    });
    fixture = TestBed.createComponent(HomeCollectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
