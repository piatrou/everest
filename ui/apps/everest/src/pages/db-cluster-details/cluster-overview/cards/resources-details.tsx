// everest
// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { DatabaseIcon, OverviewCard } from '@percona/ui-lib';
import OverviewSection from '../overview-section';
import { ResourcesDetailsOverviewProps } from './card.types';
import OverviewSectionRow from '../overview-section-row';
import { Messages } from '../cluster-overview.messages';

export const ResourcesDetails = ({
  numberOfNodes,
  cpu,
  memory,
  disk,
  loading,
}: ResourcesDetailsOverviewProps) => {
  return (
    <OverviewCard
      dataTestId="resources"
      cardHeaderProps={{
        title: Messages.titles.resources,
        avatar: <DatabaseIcon />,
        // TODO implement with EVEREST-1211
        // action: (
        //     <Button size="small" startIcon={<EditOutlinedIcon />}>
        //     Edit
        //   </Button>
        // ),
      }}
    >
      <OverviewSection
        title={`${numberOfNodes} node${+numberOfNodes > 1 ? 's' : ''}`}
        loading={loading}
      >
        <OverviewSectionRow
          label={Messages.fields.cpu}
          contentString={`${cpu}`}
        />
        <OverviewSectionRow
          label={Messages.fields.disk}
          contentString={`${disk}`}
        />
        <OverviewSectionRow
          label={Messages.fields.memory}
          contentString={`${memory}`}
        />
      </OverviewSection>
    </OverviewCard>
  );
};
