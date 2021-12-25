# T1
create view CS_View as select * from Student where Sdept = 'CS';
# T2
SELECT CS_View.* from CS_View,SC where CS_View.Sno = SC.Sno and Cno = '1';
# T3
CREATE VIEW IS_View as SELECT Student.* from Student,SC where Student.Sno = SC.Sno and Grade > 80 and Sdept = 'IS';
# T4
SELECT * from IS_View;
# T5
DROP VIEW IS_View;
# T6
create user U1,U2;
grant select, update on Student to U1;
grant insert on SC to U2;
# T7
REVOKE ALL PRIVILEGES, GRANT OPTION FROM U1;
SHOW GRANTS FOR U1;
# T9
CREATE TRIGGER T_scholar
    before update
    on SC
    for each row
    IF NEW.Grade > 95 THEN
        IF EXISTS(SELECT * from Student where Student.Sno = OLD.Sno and Student.Scholarship = '否') THEN
            UPDATE Student SET Scholarship = '是' where Student.Sno = OLD.Sno;
        end if;
    ELSE
        IF EXISTS(SELECT * from SC where SC.Sno = OLD.Sno and SC.Grade <= 95) THEN
            IF OLD.Grade > 95 THEN
                UPDATE Student SET Scholarship = '否' where Student.Sno = OLD.Sno;
            end if;
        end if;
    end if;
# T10
DROP TRIGGER T_scholar;
# T11
CREATE PROCEDURE ShowCSAVGMAXScore()
BEGIN
    SELECT AVG(Grade) as average, MAX(Grade) as maxium from SC where Sno in (SELECT Sno from CS_View);
end;
call ShowCSAVGMAXScore();
# T12
CREATE PROCEDURE ShowGrade(IN studentID VARCHAR(9))
BEGIN
    SELECT Student.Sno,Student.Sname,SC.Cno,SC.Grade from Student,SC where Student.Sno = SC.Sno and SC.Sno = studentID;
end;
set @studentID = '200215121';
call ShowGrade(@studentID);
# T13
CREATE FUNCTION showGradeF(studentID VARCHAR(9)) RETURNS TEXT
BEGIN
    DECLARE Result TEXT;
    SELECT CONCAT('[', GROUP_CONCAT(t.result), ']')
    INTO Result
    FROM (
             SELECT concat('{"Sno":"', Student.Sno, '","Sname":"', Student.Sname, '","Cno":"', SC.Cno, '","Grade":', SC.Grade, '}') as result
             from Student,
                  SC
             where Student.Sno = SC.Sno and Student.Sno = studentID) as t;
    RETURN Result;
end;
# T14
ALTER TABLE SC add constraint c_GradeRange check (Grade BETWEEN 0 and 100)