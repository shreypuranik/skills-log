-- name: CreateMember :exec
INSERT INTO member (name) 
VALUES ($1);

-- name: CreateSkill :exec
INSERT INTO skill (name) 
VALUES ($1);

-- name: CreateSkillToMemberRating :exec
INSERT INTO member_skill_rating (member_id, skill_id, rating) 
VALUES ($1, $2, $3);
